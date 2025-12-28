package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

func toKebabCase(s string) string {
	var res strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			res.WriteRune('-')
		}
		res.WriteRune(r)
	}
	return strings.ToLower(res.String())
}

func toSnakeCase(s string) string {
	var res strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			res.WriteRune('_')
		}
		res.WriteRune(r)
	}
	return strings.ToLower(res.String())
}

func ensureEntrypoint(pkgDir, tag string, tmpl *template.Template, config *Config) error {
	entrypointPath := filepath.Join(pkgDir, "entrypoint.gen.go")
	if _, err := os.Stat(entrypointPath); os.IsNotExist(err) {
		f, err := os.Create(entrypointPath)
		if err != nil {
			return fmt.Errorf("failed to create entrypoint for %s: %w", tag, err)
		}
		defer f.Close()
		data := TemplateData{
			PackageName: tag,
			Use:         strings.ReplaceAll(tag, "_", "-"),
			Short:       fmt.Sprintf("%s endpoints", tag),
			Aliases:     config.TagAliases[tag],
		}
		if err := tmpl.Execute(f, data); err != nil {
			return fmt.Errorf("failed to execute entrypoint template for %s: %w", tag, err)
		}
	}
	return nil
}

func resolveResourceType(spec *OpenAPI, schema map[string]interface{}) string {
	if schema == nil {
		return ""
	}

	if props, ok := schema["properties"].(map[string]interface{}); ok {
		if dataProp, ok := props["data"].(map[string]interface{}); ok {
			// Check if data is an array
			if items, ok := dataProp["items"].(map[string]interface{}); ok {
				return resolveResourceType(spec, resolveSchema(spec, items))
			}
			// Check if data is an object
			return resolveResourceType(spec, resolveSchema(spec, dataProp))
		}
		if typeProp, ok := props["type"].(map[string]interface{}); ok {
			resolvedTypeProp := resolveSchema(spec, typeProp)
			if enum, ok := resolvedTypeProp["enum"].([]interface{}); ok && len(enum) > 0 {
				return enum[0].(string)
			}
		}
	}
	return ""
}

func resolveSchema(spec *OpenAPI, schema map[string]interface{}) map[string]interface{} {
	if ref, ok := schema["$ref"].(string); ok {
		refName := filepath.Base(ref)
		if s, ok := spec.Components.Schemas[refName]; ok {
			return s.Schema
		}
	}
	return schema
}

func prepareTemplateData(tag, rawTag, apiTagName, method, path string, op Operation, spec *OpenAPI, config *Config) TemplateData {
	var args []string
	var argTypes []string
	use := toKebabCase(op.OperationID)

	for _, param := range op.Parameters {
		name := param.Name
		in := param.In
		required := param.Required
		schema := param.Schema

		if param.Ref != "" {
			refName := filepath.Base(param.Ref)
			if p, ok := spec.Components.Parameters[refName]; ok {
				name = p.Name
				in = p.In
				required = p.Required
				schema = p.Schema
			}
		}

		if (in == "path" || required) && in != "header" && in != "cookie" {
			args = append(args, name)
			use += fmt.Sprintf(" [%s]", name)
			argType := "string"

			resolvedSchema := schema
			if schema != nil {
				if ref, ok := schema["$ref"].(string); ok {
					refName := filepath.Base(ref)
					if s, ok := spec.Components.Schemas[refName]; ok {
						resolvedSchema = map[string]interface{}{
							"type":   s.Type,
							"format": s.Format,
						}
					}
				}
			}

			if resolvedSchema != nil {
				sType, _ := resolvedSchema["type"].(string)
				sFormat, _ := resolvedSchema["format"].(string)
				switch sType {
				case "integer":
					argType = "int64"
				case "number":
					argType = "float64"
				case "string":
					if sFormat == "uuid" {
						argType = "uuid.UUID"
					} else if sFormat == "date-time" {
						argType = "time.Time"
					}
				case "array":
					items, _ := resolvedSchema["items"].(map[string]interface{})
					if items != nil {
						if iType, ok := items["type"].(string); ok && iType == "string" {
							argType = "[]string"
						}
					}
				}
			}

			// Custom overrides
			if op.OperationID == "GetSBOM" && name == "asset_type" {
				argType = "datadogV2.AssetType"
			}
			if op.OperationID == "GetTeamSync" && name == "filter[source]" {
				argType = "datadogV2.TeamSyncAttributesSource"
			}
			if (op.OperationID == "GetCostByOrg" || op.OperationID == "GetHistoricalCostByOrg" || op.OperationID == "GetMonthlyCostAttribution") && name == "start_month" {
				argType = "time.Time"
			}
			if (op.OperationID == "GetHourlyUsage" || op.OperationID == "GetUsageApplicationSecurityMonitoring" || op.OperationID == "GetUsageLambdaTracedInvocations" || op.OperationID == "GetUsageObservabilityPipelines") && name == "start_hr" {
				argType = "time.Time"
			}

			argTypes = append(argTypes, argType)
		}
	}

	hasResponse := false
	for code, resp := range op.Responses {
		if code == "200" || code == "201" || code == "202" {
			if resp.Content != nil {
				hasResponse = true
				break
			}
		}
	}
	if op.OperationID == "GetSignalNotificationRules" || op.OperationID == "GetVulnerabilityNotificationRules" {
		hasResponse = true
	}

	hasRequestBody := op.RequestBody != nil
	requestBodyType := ""
	isOptionalParams := false
	if hasRequestBody {
		for _, content := range op.RequestBody.Content {
			if content.Schema.Ref != "" {
				requestBodyType = filepath.Base(content.Schema.Ref)
				break
			}
		}
		if requestBodyType == "" {
			requestBodyType = op.OperationID + "Request"
		}

		if optType, ok := config.OptionalParametersOperations[op.OperationID]; ok {
			isOptionalParams = true
			requestBodyType = optType
		}
	}

	resourceType := ""
	if hasResponse {
		for code, resp := range op.Responses {
			if code == "200" || code == "201" || code == "202" {
				if resp.Content != nil {
					if jsonContent, ok := resp.Content["application/json"].(map[string]interface{}); ok {
						if schema, ok := jsonContent["schema"].(map[string]interface{}); ok {
							resourceType = resolveResourceType(spec, resolveSchema(spec, schema))
						}
					}
				}
			}
		}
	}
	// Fallback to package name if we can't find it, or use special logic
	if resourceType == "" {
		resourceType = tag
	}

	var aliases []string
	tagKebab := strings.ReplaceAll(tag, "_", "-")
	opKebab := toKebabCase(op.OperationID)

	tagSingular := strings.TrimSuffix(tagKebab, "s")
	variants := []string{tagKebab, tagSingular}

	for _, variant := range variants {
		if variant == "" {
			continue
		}
		if strings.Contains(opKebab, variant) {
			alias := strings.ReplaceAll(opKebab, variant, "")
			alias = strings.ReplaceAll(alias, "--", "-")
			alias = strings.Trim(alias, "-")
			if alias != "" {
				aliases = append(aliases, alias)
				break
			}
		}
	}

	return TemplateData{
		PackageName:      tag,
		CommandName:      strings.ToUpper(op.OperationID[:1]) + op.OperationID[1:] + "Cmd",
		Use:              use,
		Short:            op.Summary,
		Method:           strings.ToUpper(method),
		Path:             path,
		OperationID:      op.OperationID,
		TagName:          rawTag,
		ApiTagName:       apiTagName,
		Args:             args,
		ArgTypes:         argTypes,
		HasRequestBody:   hasRequestBody,
		RequestBodyType:  requestBodyType,
		IsOptionalParams: isOptionalParams,
		HasResponse:      hasResponse,
		ResourceType:     resourceType,
		Aliases:          aliases,
	}
}

func generateCommandFile(pkgDir, operationID string, data TemplateData, tmpl *template.Template) error {
	fileName := fmt.Sprintf("%s.gen.go", toSnakeCase(operationID))
	filePath := filepath.Join(pkgDir, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template for %s: %w", filePath, err)
	}
	fmt.Printf("Generated %s\n", filePath)
	return nil
}

func updateRootGo(tags map[string]bool) error {
	var sortedTags []string
	for tag := range tags {
		sortedTags = append(sortedTags, tag)
	}
	sort.Strings(sortedTags)

	var imports []string
	var commands []string
	for _, tag := range sortedTags {
		imports = append(imports, fmt.Sprintf("ouaf/cmd/%s", tag))
		commands = append(commands, fmt.Sprintf("%s.Cmd", tag))
	}

	tmpl, err := template.ParseFiles("generator/root.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse root template: %w", err)
	}

	f, err := os.Create("cmd/root.gen.go")
	if err != nil {
		return fmt.Errorf("failed to create cmd/root.gen.go: %w", err)
	}
	defer f.Close()

	data := RootTemplateData{
		Imports:  imports,
		Commands: commands,
	}

	return tmpl.Execute(f, data)
}
