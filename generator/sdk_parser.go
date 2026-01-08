package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// parseSDK parses the Datadog Go SDK and extracts operation information
func parseSDK() (*SDKSpec, error) {
	// Get the SDK path from go modules
	sdkPath, err := getSDKPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get SDK path: %w", err)
	}

	apiDir := filepath.Join(sdkPath, "api", "datadogV2")
	if _, err := os.Stat(apiDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("SDK API directory not found: %s", apiDir)
	}

	spec := &SDKSpec{
		Operations: make(map[string]*SDKOperation),
	}

	// Parse all API files
	files, err := filepath.Glob(filepath.Join(apiDir, "api_*.go"))
	if err != nil {
		return nil, fmt.Errorf("failed to list API files: %w", err)
	}

	for _, file := range files {
		if err := parseAPIFile(file, spec); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", file, err)
		}
	}

	return spec, nil
}

// getSDKPath returns the path to the Datadog Go SDK
func getSDKPath() (string, error) {
	// Use go list to find the module path
	modPath := os.Getenv("GOMODCACHE")
	if modPath == "" {
		modPath = filepath.Join(os.Getenv("HOME"), "go", "pkg", "mod")
	}

	// The module is at: github.com/!data!Dog/datadog-api-client-go/v2@v2.52.0
	// We need to find the exact version from go.mod
	version := "v2.52.0" // TODO: Extract from go.mod or use go list
	sdkPath := filepath.Join(modPath, "github.com", "!data!Dog", "datadog-api-client-go", "v2@"+version)

	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		return "", fmt.Errorf("SDK not found at %s", sdkPath)
	}

	return sdkPath, nil
}

// parseAPIFile parses a single API file and extracts operations
func parseAPIFile(filename string, spec *SDKSpec) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Extract the API name from the filename (e.g., api_action_connection.go -> action_connection)
	baseName := filepath.Base(filename)
	apiName := strings.TrimPrefix(baseName, "api_")
	apiName = strings.TrimSuffix(apiName, ".go")

	// Find the API type (e.g., ActionConnectionApi)
	var apiType string
	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if strings.HasSuffix(typeSpec.Name.Name, "Api") {
				apiType = typeSpec.Name.Name
				return false
			}
		}
		return true
	})

	if apiType == "" {
		// Skip files that don't define an API type
		return nil
	}

	// Extract the tag name (e.g., ActionConnectionApi -> Action Connection)
	tag := extractTagFromAPIType(apiType)

	// Parse all methods of this API
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
			continue
		}

		// Check if this is a method of the API type
		recvType := getReceiverType(funcDecl.Recv)
		if recvType != apiType {
			continue
		}

		// Parse the method
		op, err := parseMethod(funcDecl, tag, apiName, fset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to parse method %s: %v\n", funcDecl.Name.Name, err)
			continue
		}

		if op != nil {
			spec.Operations[op.OperationID] = op
		}
	}

	return nil
}

// getReceiverType extracts the receiver type name from a method
func getReceiverType(recv *ast.FieldList) string {
	if len(recv.List) == 0 {
		return ""
	}

	field := recv.List[0]
	switch t := field.Type.(type) {
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	case *ast.Ident:
		return t.Name
	}

	return ""
}

// parseMethod parses a method declaration and extracts operation information
func parseMethod(funcDecl *ast.FuncDecl, tag, apiName string, fset *token.FileSet) (*SDKOperation, error) {
	op := &SDKOperation{
		OperationID: funcDecl.Name.Name,
		Tags:        []string{tag},
		APIName:     apiName,
		Parameters:  []SDKParameter{},
	}

	// Extract summary and description from comments
	if funcDecl.Doc != nil {
		op.Summary, op.Description = extractDocumentation(funcDecl.Doc.Text())
	}

	// Parse function parameters
	if funcDecl.Type.Params != nil {
		for _, param := range funcDecl.Type.Params.List {
			// Skip context parameter
			if len(param.Names) > 0 && param.Names[0].Name == "ctx" {
				continue
			}

			// Check if this is an optional parameters struct
			paramType := getTypeString(param.Type)
			if strings.HasPrefix(paramType, "...") && strings.HasSuffix(paramType, "OptionalParameters") {
				// This will be parsed from the method body
				op.HasOptionalParams = true
				continue
			}

			// Regular parameter (path or body)
			for _, name := range param.Names {
				sdkParam := SDKParameter{
					Name:     name.Name,
					Type:     paramType,
					Required: true,
				}

				// Determine if it's a body parameter
				if name.Name == "body" {
					sdkParam.In = "body"
					op.RequestBody = &SDKRequestBody{
						Required: true,
						Type:     paramType,
					}
				} else {
					// Assume it's a path parameter; we'll confirm when parsing the method body
					sdkParam.In = "path"
				}

				op.Parameters = append(op.Parameters, sdkParam)
			}
		}
	}

	// Check if this is a pagination helper (returns channel and cancel func)
	// These methods have names like "ListUsersWithPagination" and have different signatures
	if strings.HasSuffix(op.OperationID, "WithPagination") {
		// Skip pagination helpers - they're not standard API operations
		return nil, nil
	}

	// Parse return type
	if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) > 0 {
		firstResult := funcDecl.Type.Results.List[0]
		resultType := getTypeString(firstResult.Type)

		// If the first return value is *http.Response (mapped from *_nethttp.Response),
		// it means there's no response body (delete operations, etc.)
		if resultType == "*http.Response" {
			op.HasResponse = false
		} else {
			// Otherwise, this is the actual response type
			op.ResponseType = resultType
			op.HasResponse = true
		}
	}

	// Parse method body to extract HTTP method and path
	if funcDecl.Body != nil {
		parseMethodBody(funcDecl.Body, op)
	}

	// Mark non-body parameters that are used in path replacement as path parameters
	// The SDK uses snake_case placeholders in paths but camelCase in parameter names
	if op.Path != "" {
		for i := range op.Parameters {
			if op.Parameters[i].In != "body" {
				// Convert camelCase parameter name to snake_case for matching
				snakeCaseName := toSnakeCaseSimple(op.Parameters[i].Name)
				paramPlaceholder := "{" + snakeCaseName + "}"
				if strings.Contains(op.Path, paramPlaceholder) {
					op.Parameters[i].In = "path"
				}
			}
		}
	}

	return op, nil
}

// parseMethodBody parses the method body to extract HTTP method, path, and parameter details
func parseMethodBody(body *ast.BlockStmt, op *SDKOperation) {
	ast.Inspect(body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			// Look for localVarHTTPMethod = _nethttp.MethodXXX
			if len(node.Lhs) == 1 && len(node.Rhs) == 1 {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarHTTPMethod" {
					if sel, ok := node.Rhs[0].(*ast.SelectorExpr); ok {
						op.Method = strings.TrimPrefix(sel.Sel.Name, "Method")
					}
				}
			}

			// Look for localVarPostBody = &body
			if len(node.Lhs) == 1 && len(node.Rhs) == 1 {
				if ident, ok := node.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarPostBody" {
					// Confirms this method has a request body
					if op.RequestBody != nil {
						op.RequestBody.Required = true
					}
				}
			}

		case *ast.ValueSpec:
			// Look for localVarPath := localBasePath + "/api/v2/..."
			if len(node.Names) == 1 && node.Names[0].Name == "localVarPath" {
				if len(node.Values) > 0 {
					if binExpr, ok := node.Values[0].(*ast.BinaryExpr); ok && binExpr.Op == token.ADD {
						if lit, ok := binExpr.Y.(*ast.BasicLit); ok {
							op.Path = lit.Value[1 : len(lit.Value)-1] // Remove quotes
						}
					}
				}
			}
		}
		return true
	})
}

// extractTagFromAPIType converts API type name to tag (e.g., ActionConnectionApi -> Action Connection)
func extractTagFromAPIType(apiType string) string {
	// Remove "Api" suffix
	name := strings.TrimSuffix(apiType, "Api")

	// Insert spaces before capital letters
	var result strings.Builder
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}

	return result.String()
}

// extractDocumentation extracts summary and description from doc comment
func extractDocumentation(docText string) (summary, description string) {
	lines := strings.Split(strings.TrimSpace(docText), "\n")
	if len(lines) == 0 {
		return "", ""
	}

	// First line is: MethodName Summary text
	firstLine := lines[0]
	parts := strings.SplitN(firstLine, " ", 2)
	if len(parts) > 1 {
		summary = strings.TrimSpace(parts[1])
	}

	// Remaining lines are description
	if len(lines) > 1 {
		description = strings.TrimSpace(strings.Join(lines[1:], "\n"))
	}

	return summary, description
}

// getTypeString converts an AST type expression to a string
func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	case *ast.SelectorExpr:
		pkgName := getTypeString(t.X)
		// Map SDK package aliases to standard package names
		pkgName = mapSDKPackageAlias(pkgName)
		return pkgName + "." + t.Sel.Name
	case *ast.Ellipsis:
		return "..." + getTypeString(t.Elt)
	default:
		return ""
	}
}

// mapSDKPackageAlias maps SDK package aliases to their standard names
// E.g., "_io" -> "io", "_nethttp" -> "http", "_context" -> "context"
// Note: We use "http" not "net/http" because "/" breaks Go syntax in templates
func mapSDKPackageAlias(pkgAlias string) string {
	aliases := map[string]string{
		"_io":      "io",
		"_nethttp": "http", // Use http, templates will import "net/http" as http
		"_context": "context",
		"_neturl":  "url", // Use url, templates will import "net/url" as url
		"_fmt":     "fmt",
		"_time":    "time",
	}

	if mapped, ok := aliases[pkgAlias]; ok {
		return mapped
	}
	return pkgAlias
}

// toSnakeCaseSimple converts camelCase to snake_case
func toSnakeCaseSimple(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
