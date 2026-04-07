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

		if err := os.WriteFile(entrypointPath, formatted, 0o644); err != nil {
			return fmt.Errorf("failed to write entrypoint for %s: %w", bundle, err)
		}
	}
	return nil
}

func prepareTemplateData(op OperationModel, config *Config) TemplateData {
	args := make([]string, 0, len(op.Parameters))
	argTypes := make([]string, 0, len(op.Parameters))
	use := toKebabCase(op.OperationID, config)

	for _, param := range op.Parameters {
		args = append(args, param.Name)
		use += fmt.Sprintf(" [%s]", param.Name)
		argTypes = append(argTypes, normalizeArgumentType(param.GoType))
	}

	resourceType := op.ResourceType
	if resourceType == "" {
		resourceType = op.Bundle
	}

	responseTypeGo := op.ResponseTypeGo
	if responseTypeGo == "" {
		responseTypeGo = "interface{}"
	}

	aliases := computeAliases(op.Bundle, op.OperationID, config)
	docBundle := strings.ReplaceAll(op.Bundle, "_", "-")
	docURL := fmt.Sprintf("https://docs.datadoghq.com/api/latest/%s/#%s", docBundle, toKebabCase(op.OperationID, config))

	return TemplateData{
		PackageName:      op.Bundle,
		CommandName:      strings.ToUpper(op.OperationID[:1]) + op.OperationID[1:] + "Cmd",
		Use:              use,
		Short:            strings.TrimSuffix(op.Summary, "."),
		Method:           strings.ToUpper(op.Method),
		Path:             op.Path,
		OperationID:      op.OperationID,
		BundleName:       op.Bundle,
		ApiBundleName:    op.ApiBundleName,
		Args:             args,
		ArgTypes:         argTypes,
		HasRequestBody:   op.HasRequestBody,
		RequestBodyType:  op.RequestBodyType,
		IsOptionalParams: op.IsOptionalParams,
		HasResponse:      op.HasResponse,
		ResourceType:     resourceType,
		ResponseTypeGo:   responseTypeGo,
		Aliases:          aliases,
		DocURL:           docURL,
	}
}

func normalizeArgumentType(goType string) string {
	switch goType {
	case "uuid.UUID", "time.Time", "int", "int32", "int64", "float64", "[]string":
		return goType
	}
	if looksLikeSDKAliasType(goType) {
		return "datadogV2." + goType
	}
	return "string"
}

func looksLikeSDKAliasType(goType string) bool {
	if goType == "" {
		return false
	}
	if strings.Contains(goType, ".") || strings.Contains(goType, "[]") || strings.Contains(goType, "*") {
		return false
	}
	switch goType {
	case "string", "bool", "int", "int32", "int64", "float32", "float64":
		return false
	}
	r := rune(goType[0])
	return r >= 'A' && r <= 'Z'
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

	if err := os.WriteFile(filePath, formatted, 0o644); err != nil {
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
		imports = append(imports, fmt.Sprintf("github.com/gplassard/woof/cmd/%s", bundle))
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

	return os.WriteFile("cmd/root.gen.go", formatted, 0o644)
}
