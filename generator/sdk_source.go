package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func loadFromSDK(config *Config) (*GenerationInput, error) {
	sdkDir, err := sdkPackageDir()
	if err != nil {
		return nil, err
	}

	sdkModel, err := parseSDKModel(sdkDir, config)
	if err != nil {
		return nil, err
	}

	return &GenerationInput{Operations: sdkModel.Operations}, nil
}

func parseSDKModel(sdkDir string, config *Config) (*SDKModel, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, sdkDir, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sdk directory: %w", err)
	}
	pkg, ok := pkgs["datadogV2"]
	if !ok {
		return nil, fmt.Errorf("datadogV2 package not found in %s", sdkDir)
	}

	model := &SDKModel{
		Structs: map[string]SDKStruct{},
		Enums:   map[string]SDKEnum{},
	}

	files := make([]*ast.File, 0, len(pkg.Files))
	for _, f := range pkg.Files {
		files = append(files, f)
	}

	for _, f := range files {
		collectStructsAndEnums(model, f)
	}

	services := map[string]*SDKService{}
	operations := make([]OperationModel, 0, 1024)
	for _, f := range files {
		if !strings.HasPrefix(filepath.Base(fset.Position(f.Pos()).Filename), "api_") {
			continue
		}
		for _, decl := range f.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv == nil || fn.Name == nil || fn.Body == nil {
				continue
			}
			recvName := receiverTypeName(fn)
			if !strings.HasSuffix(recvName, "Api") {
				continue
			}
			if strings.HasPrefix(fn.Name.Name, "New") || strings.HasSuffix(fn.Name.Name, "WithPagination") {
				continue
			}
			op, ok := parseOperation(fn, recvName, model, config)
			if !ok {
				continue
			}
			operations = append(operations, op)
			svc := services[recvName]
			if svc == nil {
				svc = &SDKService{Name: recvName, Bundle: op.Bundle, ApiName: op.ApiBundleName}
				services[recvName] = svc
			}
			svc.Operations = append(svc.Operations, op.OperationID)
		}
	}

	sort.Slice(operations, func(i, j int) bool {
		if operations[i].Bundle == operations[j].Bundle {
			return operations[i].OperationID < operations[j].OperationID
		}
		return operations[i].Bundle < operations[j].Bundle
	})
	model.Operations = operations

	serviceNames := make([]string, 0, len(services))
	for name := range services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)
	for _, name := range serviceNames {
		svc := services[name]
		sort.Strings(svc.Operations)
		model.Services = append(model.Services, *svc)
	}

	return model, nil
}

func collectStructsAndEnums(model *SDKModel, f *ast.File) {
	for _, decl := range f.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		switch gen.Tok {
		case token.TYPE:
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				fields := map[string]SDKField{}
				for _, field := range st.Fields.List {
					if len(field.Names) == 0 {
						continue
					}
					jsonName := ""
					if field.Tag != nil {
						tagValue, err := strconv.Unquote(field.Tag.Value)
						if err == nil {
							jsonName = strings.Split(reflect.StructTag(tagValue).Get("json"), ",")[0]
						}
					}
					if jsonName == "" || jsonName == "-" {
						continue
					}
					typeName := unwrapTypeName(field.Type)
					if typeName == "" {
						continue
					}
					fields[jsonName] = SDKField{JSONName: jsonName, TypeName: typeName}
				}
				model.Structs[ts.Name.Name] = SDKStruct{Name: ts.Name.Name, Fields: fields}
			}
		case token.CONST:
			for _, spec := range gen.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				typeIdent, ok := vs.Type.(*ast.Ident)
				if !ok {
					continue
				}
				enum := model.Enums[typeIdent.Name]
				enum.Name = typeIdent.Name
				for _, value := range vs.Values {
					lit, ok := value.(*ast.BasicLit)
					if !ok || lit.Kind != token.STRING {
						continue
					}
					v, err := strconv.Unquote(lit.Value)
					if err != nil {
						continue
					}
					enum.Values = append(enum.Values, v)
				}
				if len(enum.Values) > 0 {
					model.Enums[typeIdent.Name] = enum
				}
			}
		}
	}
}

func parseOperation(fn *ast.FuncDecl, recvName string, model *SDKModel, config *Config) (OperationModel, bool) {
	operationID := fn.Name.Name
	apiName := strings.TrimSuffix(recvName, "Api")
	bundle := normalizeBundleFromSDK(apiName, config)

	httpMethod := ""
	path := ""
	pathParamNames := map[string]string{}
	queryParamNames := map[string]string{}
	bodyParam := ""
	analyzeFunctionBody(fn.Body, &httpMethod, &path, pathParamNames, queryParamNames, &bodyParam)

	if operationID == "" || httpMethod == "" || path == "" {
		return OperationModel{}, false
	}

	params := extractParameters(fn, bodyParam, pathParamNames, queryParamNames)
	responseTypeGo, hasResponse := extractResponseType(fn)

	hasRequestBody := bodyParam != ""
	requestBodyType := ""
	if hasRequestBody {
		requestBodyType = findParameterType(fn, bodyParam)
	}
	isOptionalParams := false
	if optType, ok := config.OptionalParametersOperations[operationID]; ok {
		hasRequestBody = true
		isOptionalParams = true
		requestBodyType = optType
	}
	if hasRequestBody && requestBodyType == "" {
		requestBodyType = operationID + "Request"
	}

	resourceType := ""
	if hasResponse {
		resourceType = inferResourceType(model, responseTypeGo)
	}

	return OperationModel{
		OperationID:      operationID,
		Summary:          extractSummary(fn, operationID, config),
		Method:           strings.TrimPrefix(httpMethod, "Method"),
		Path:             path,
		Bundle:           bundle,
		ApiBundleName:    apiName,
		Parameters:       params,
		HasRequestBody:   hasRequestBody,
		RequestBodyType:  requestBodyType,
		IsOptionalParams: isOptionalParams,
		HasResponse:      hasResponse,
		ResponseTypeGo:   responseTypeGo,
		ResourceType:     resourceType,
	}, true
}

func receiverTypeName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}
	switch t := fn.Recv.List[0].Type.(type) {
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	case *ast.Ident:
		return t.Name
	}
	return ""
}

func analyzeFunctionBody(body *ast.BlockStmt, httpMethod *string, path *string, pathParamNames map[string]string, queryParamNames map[string]string, bodyParam *string) {
	ast.Inspect(body, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ValueSpec:
			for i, name := range x.Names {
				if name.Name == "localVarHTTPMethod" && i < len(x.Values) {
					if sel, ok := x.Values[i].(*ast.SelectorExpr); ok {
						*httpMethod = sel.Sel.Name
					}
				}
			}
		case *ast.AssignStmt:
			for i, lhs := range x.Lhs {
				ident, ok := lhs.(*ast.Ident)
				if !ok || i >= len(x.Rhs) {
					continue
				}
				rhs := x.Rhs[i]
				if ident.Name == "localVarPath" {
					if p := extractPathLiteral(rhs); p != "" {
						*path = p
					}
					if call, ok := rhs.(*ast.CallExpr); ok {
						extractPathParamMapping(call, pathParamNames)
					}
				}
				if ident.Name == "localVarPostBody" {
					if unary, ok := rhs.(*ast.UnaryExpr); ok {
						if arg, ok := unary.X.(*ast.Ident); ok {
							*bodyParam = arg.Name
						}
					}
				}
			}
		case *ast.CallExpr:
			extractQueryParamMapping(x, queryParamNames)
		}
		return true
	})
}

func extractPathLiteral(expr ast.Expr) string {
	if be, ok := expr.(*ast.BinaryExpr); ok && be.Op == token.ADD {
		if lit, ok := be.Y.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			v, err := strconv.Unquote(lit.Value)
			if err == nil {
				return v
			}
		}
	}
	return ""
}

func extractPathParamMapping(call *ast.CallExpr, pathParamNames map[string]string) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "ReplacePathParameter" || len(call.Args) < 3 {
		return
	}
	placeholderLit, ok := call.Args[1].(*ast.BasicLit)
	if !ok || placeholderLit.Kind != token.STRING {
		return
	}
	placeholder, err := strconv.Unquote(placeholderLit.Value)
	if err != nil {
		return
	}
	placeholder = strings.Trim(placeholder, "{}")

	parameterName := extractParameterToStringIdent(call.Args[2])
	if parameterName != "" {
		pathParamNames[parameterName] = placeholder
	}
}

func extractQueryParamMapping(call *ast.CallExpr, queryParamNames map[string]string) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || (sel.Sel.Name != "Add" && sel.Sel.Name != "Set") || len(call.Args) < 2 {
		return
	}
	xIdent, ok := sel.X.(*ast.Ident)
	if !ok || xIdent.Name != "localVarQueryParams" {
		return
	}
	queryNameLit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || queryNameLit.Kind != token.STRING {
		return
	}
	queryName, err := strconv.Unquote(queryNameLit.Value)
	if err != nil {
		return
	}
	parameterName := extractParameterToStringIdent(call.Args[1])
	if parameterName != "" {
		queryParamNames[parameterName] = queryName
	}
}

func extractParameterToStringIdent(expr ast.Expr) string {
	call, ok := expr.(*ast.CallExpr)
	if !ok || len(call.Args) == 0 {
		return ""
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "ParameterToString" {
		return ""
	}
	switch arg := call.Args[0].(type) {
	case *ast.Ident:
		return arg.Name
	case *ast.StarExpr:
		if ident, ok := arg.X.(*ast.SelectorExpr); ok {
			if id, ok := ident.X.(*ast.Ident); ok && id.Name == "optionalParams" {
				return ""
			}
		}
	}
	return ""
}

func extractParameters(fn *ast.FuncDecl, bodyParam string, pathParamNames map[string]string, queryParamNames map[string]string) []ParameterModel {
	params := make([]ParameterModel, 0)
	if fn.Type.Params == nil {
		return params
	}

	for _, field := range fn.Type.Params.List {
		if len(field.Names) == 0 {
			continue
		}
		if _, ok := field.Type.(*ast.Ellipsis); ok {
			continue
		}
		for _, name := range field.Names {
			if name.Name == "ctx" || name.Name == bodyParam {
				continue
			}
			param := ParameterModel{
				GoName: name.Name,
				GoType: normalizeGoType(exprString(field.Type)),
				In:     "query",
			}
			if p, ok := pathParamNames[name.Name]; ok {
				param.Name = p
				param.In = "path"
			} else if q, ok := queryParamNames[name.Name]; ok {
				param.Name = q
				param.In = "query"
			} else {
				param.Name = toSnakeCase(name.Name, nil)
			}
			params = append(params, param)
		}
	}
	return params
}

func normalizeBundleFromSDK(apiName string, config *Config) string {
	if config != nil && config.BundleNameOverrides != nil {
		if mapped, ok := config.BundleNameOverrides[apiName]; ok {
			return mapped
		}
	}
	return toSnakeCase(apiName, config)
}

func extractResponseType(fn *ast.FuncDecl) (string, bool) {
	if fn.Type.Results == nil || len(fn.Type.Results.List) == 0 {
		return "", false
	}
	first := exprString(fn.Type.Results.List[0].Type)
	if first == "*_nethttp.Response" || first == "*nethttp.Response" {
		return "", false
	}
	if strings.HasSuffix(first, ".Reader") || first == "_io.Reader" || first == "io.Reader" {
		return "interface{}", true
	}
	builtin := map[string]bool{
		"string":                 true,
		"bool":                   true,
		"int":                    true,
		"int32":                  true,
		"int64":                  true,
		"float32":                true,
		"float64":                true,
		"[]byte":                 true,
		"map[string]interface{}": true,
	}
	if builtin[first] {
		return "interface{}", true
	}
	if strings.HasPrefix(first, "*") {
		first = strings.TrimPrefix(first, "*")
	}
	if strings.Contains(first, ".") {
		parts := strings.Split(first, ".")
		first = parts[len(parts)-1]
	}
	return first, true
}

func findParameterType(fn *ast.FuncDecl, paramName string) string {
	if fn.Type.Params == nil {
		return ""
	}
	for _, field := range fn.Type.Params.List {
		for _, name := range field.Names {
			if name.Name == paramName {
				t := normalizeGoType(exprString(field.Type))
				if strings.Contains(t, ".") {
					parts := strings.Split(t, ".")
					t = parts[len(parts)-1]
				}
				return strings.TrimPrefix(t, "*")
			}
		}
	}
	return ""
}

func extractSummary(fn *ast.FuncDecl, operationID string, config *Config) string {
	if fn.Doc == nil {
		return humanizeOperationID(operationID, config)
	}
	doc := strings.TrimSpace(fn.Doc.Text())
	if doc == "" {
		return humanizeOperationID(operationID, config)
	}
	line := strings.Split(doc, "\n")[0]
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, operationID+" ")
	line = strings.TrimSpace(line)
	if line == "" || line == operationID {
		return humanizeOperationID(operationID, config)
	}
	return line
}

func humanizeOperationID(operationID string, config *Config) string {
	phrase := strings.ReplaceAll(toKebabCase(operationID, config), "-", " ")
	if phrase == "" {
		return operationID
	}
	return strings.ToUpper(phrase[:1]) + phrase[1:]
}

func exprString(expr ast.Expr) string {
	var b strings.Builder
	_ = printer.Fprint(&b, token.NewFileSet(), expr)
	return b.String()
}

func normalizeGoType(t string) string {
	t = strings.TrimPrefix(t, "*")
	if strings.HasPrefix(t, "[]") {
		inner := normalizeGoType(strings.TrimPrefix(t, "[]"))
		if inner == "string" {
			return "[]string"
		}
		return "[]" + inner
	}
	if strings.Contains(t, ".") {
		parts := strings.Split(t, ".")
		if len(parts) == 2 {
			return parts[0] + "." + parts[1]
		}
	}
	return t
}

func inferResourceType(model *SDKModel, responseTypeGo string) string {
	if responseTypeGo == "" || responseTypeGo == "interface{}" {
		return ""
	}
	visited := map[string]bool{}
	var walkType func(typeName string) string
	walkType = func(typeName string) string {
		typeName = strings.TrimPrefix(typeName, "*")
		typeName = strings.TrimPrefix(typeName, "[]")
		if visited[typeName] {
			return ""
		}
		visited[typeName] = true

		if enum, ok := model.Enums[typeName]; ok && len(enum.Values) > 0 {
			return enum.Values[0]
		}
		st, ok := model.Structs[typeName]
		if !ok {
			return ""
		}
		if dataField, ok := st.Fields["data"]; ok {
			if rt := walkType(dataField.TypeName); rt != "" {
				return rt
			}
		}
		if typeField, ok := st.Fields["type"]; ok {
			if rt := walkType(typeField.TypeName); rt != "" {
				return rt
			}
		}
		return ""
	}
	return walkType(responseTypeGo)
}

func unwrapTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return unwrapTypeName(t.X)
	case *ast.ArrayType:
		return "[]" + unwrapTypeName(t.Elt)
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
	}
	return ""
}
