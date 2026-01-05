package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

func toKebabCase(s string, config *Config) string {
	var acronyms []string
	if config != nil {
		acronyms = config.Acronyms
	}

	res := s
	for _, acronym := range acronyms {
		// Replace acronym with its lowercase version surrounded by markers to avoid further splitting
		// e.g. CSM -> -csm-
		// We use a temporary placeholder to avoid double replacement or splitting
		res = strings.ReplaceAll(res, acronym, "-"+strings.ToLower(acronym)+"-")
	}

	var finalRes strings.Builder
	for i, r := range res {
		if i > 0 && r >= 'A' && r <= 'Z' {
			finalRes.WriteRune('-')
		}
		finalRes.WriteRune(r)
	}

	s = strings.ToLower(finalRes.String())
	// Clean up double hyphens and hyphens at start/end
	s = strings.ReplaceAll(s, "--", "-")
	s = strings.Trim(s, "-")
	return s
}

func toSnakeCase(s string, config *Config) string {
	var acronyms []string
	if config != nil {
		acronyms = config.Acronyms
	}

	res := s
	for _, acronym := range acronyms {
		res = strings.ReplaceAll(res, acronym, "_"+strings.ToLower(acronym)+"_")
	}

	var finalRes strings.Builder
	for i, r := range res {
		if i > 0 && r >= 'A' && r <= 'Z' {
			finalRes.WriteRune('_')
		}
		finalRes.WriteRune(r)
	}

	s = strings.ToLower(finalRes.String())
	s = strings.ReplaceAll(s, "__", "_")
	s = strings.Trim(s, "_")
	return s
}

func ensureEntrypoint(pkgDir, bundle string, tmpl *template.Template, config *Config) error {
	entrypointPath := filepath.Join(pkgDir, "entrypoint.gen.go")
	if _, err := os.Stat(entrypointPath); os.IsNotExist(err) {
		data := TemplateData{
			PackageName: bundle,
			Use:         strings.ReplaceAll(bundle, "_", "-"),
			Short:       fmt.Sprintf("%s endpoints", bundle),
			Aliases:     config.BundleAliases[bundle],
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return fmt.Errorf("failed to execute entrypoint template for %s: %w", bundle, err)
		}

		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			return fmt.Errorf("failed to format entrypoint for %s: %w", bundle, err)
		}

		if err := os.WriteFile(entrypointPath, formatted, 0644); err != nil {
			return fmt.Errorf("failed to write entrypoint for %s: %w", bundle, err)
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

func prepareTemplateData(bundle, rawBundle, apiBundleName, method, path string, op Operation, spec *OpenAPI, config *Config) TemplateData {
	var args []string
	var argTypes []string
	use := toKebabCase(op.OperationID, config)

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
		resourceType = bundle
	}

	aliases := computeAliases(bundle, op.OperationID, config)

	return TemplateData{
		PackageName:      bundle,
		CommandName:      strings.ToUpper(op.OperationID[:1]) + op.OperationID[1:] + "Cmd",
		Use:              use,
		Short:            op.Summary,
		Method:           strings.ToUpper(method),
		Path:             path,
		OperationID:      op.OperationID,
		BundleName:       rawBundle,
		ApiBundleName:    apiBundleName,
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

func computeAliases(bundle, operationID string, config *Config) []string {
	var aliases []string
	bundleKebab := strings.ReplaceAll(bundle, "_", "-")
	opKebab := toKebabCase(operationID, config)

	bundleSingular := strings.TrimSuffix(bundleKebab, "s")
	bundlePlural := bundleKebab
	if !strings.HasSuffix(bundleKebab, "s") {
		bundlePlural = bundleKebab + "s"
	}
	variants := []string{bundleKebab, bundleSingular, bundlePlural}
	if config != nil {
		if bundleAliases, ok := config.BundleAliases[bundle]; ok {
			for _, alias := range bundleAliases {
				variants = append(variants, strings.ReplaceAll(alias, "_", "-"))
			}
		}
	}

	for _, variant := range variants {
		if variant == "" {
			continue
		}
		// Only replace if it matches a whole segment to avoid "list-spans-get" -> "list-s-get"
		parts := strings.Split(opKebab, "-")
		variantParts := strings.Split(variant, "-")

		var newParts []string
		skip := 0
		for i := 0; i < len(parts); i++ {
			if skip > 0 {
				skip--
				continue
			}

			match := true
			if i+len(variantParts) <= len(parts) {
				for j := 0; j < len(variantParts); j++ {
					if parts[i+j] != variantParts[j] {
						match = false
						break
					}
				}
			} else {
				match = false
			}

			if match {
				skip = len(variantParts) - 1
			} else {
				newParts = append(newParts, parts[i])
			}
		}

		alias := strings.Join(newParts, "-")
		if alias != "" && alias != opKebab {
			aliases = append(aliases, alias)
		}
	}

	// Keep only the shortest alias
	if len(aliases) > 1 {
		sort.Slice(aliases, func(i, j int) bool {
			return len(aliases[i]) < len(aliases[j])
		})
		uniqueAliases := []string{aliases[0]}
		for i := 1; i < len(aliases); i++ {
			isRedundant := false
			for _, keep := range uniqueAliases {
				if aliases[i] == keep {
					isRedundant = true
					break
				}
			}
			if !isRedundant {
				uniqueAliases = append(uniqueAliases, aliases[i])
			}
		}
		aliases = uniqueAliases
	}

	// If we have multiple aliases, keep only the shortest ones (don't keep both 'create' and 'create-action')
	if len(aliases) > 1 {
		sort.Slice(aliases, func(i, j int) bool {
			return len(aliases[i]) < len(aliases[j])
		})
		var finalAliases []string
		for _, a := range aliases {
			isSubset := false
			for _, other := range finalAliases {
				if strings.Contains(a, other) {
					isSubset = true
					break
				}
			}
			if !isSubset {
				finalAliases = append(finalAliases, a)
			}
		}
		aliases = finalAliases
	}
	return aliases
}

func generateCommandFile(pkgDir, operationID string, data TemplateData, tmpl *template.Template, config *Config) error {
	fileName := fmt.Sprintf("%s.gen.go", toSnakeCase(operationID, config))
	filePath := filepath.Join(pkgDir, fileName)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute template for %s: %w", filePath, err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format %s: %w", filePath, err)
	}

	if err := os.WriteFile(filePath, formatted, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", filePath, err)
	}

	fmt.Printf("Generated %s\n", filePath)
	return nil
}

func updateRootGo(bundles map[string]bool) error {
	var sortedBundles []string
	for bundle := range bundles {
		sortedBundles = append(sortedBundles, bundle)
	}
	sort.Strings(sortedBundles)

	var imports []string
	var commands []string
	for _, bundle := range sortedBundles {
		imports = append(imports, fmt.Sprintf("github.com/gplassard/ouaf/cmd/%s", bundle))
		commands = append(commands, fmt.Sprintf("%s.Cmd", bundle))
	}

	rootTmplContent, err := generatorFS.ReadFile("root.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to read root template: %w", err)
	}
	tmpl, err := template.New("root.go.tmpl").Parse(string(rootTmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse root template: %w", err)
	}

	data := RootTemplateData{
		Imports:  imports,
		Commands: commands,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to execute root template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format cmd/root.gen.go: %w", err)
	}

	return os.WriteFile("cmd/root.gen.go", formatted, 0644)
}
