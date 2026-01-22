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

	// Replace spaces with hyphens first
	s = strings.ReplaceAll(s, " ", "-")

	// Convert to kebab-case
	var finalRes strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			finalRes.WriteRune('-')
		}
		finalRes.WriteRune(r)
	}

	result := strings.ToLower(finalRes.String())

	// Clean up multiple hyphens first
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}

	// Then apply acronym replacements on the cleaned kebab-case string
	for _, acronym := range acronyms {
		lowerAcronym := strings.ToLower(acronym)
		// Replace the acronym with hyphens (e.g., "a-p-i" -> "api")
		parts := strings.Split(lowerAcronym, "")
		withHyphens := strings.Join(parts, "-")
		result = strings.ReplaceAll(result, withHyphens, lowerAcronym)
	}

	// Trim any leading/trailing hyphens
	result = strings.Trim(result, "-")
	return result
}

func toSnakeCase(s string, config *Config) string {
	var acronyms []string
	if config != nil {
		acronyms = config.Acronyms
	}

	// Replace spaces with underscores first
	s = strings.ReplaceAll(s, " ", "_")

	// Convert to snake_case
	var finalRes strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			finalRes.WriteRune('_')
		}
		finalRes.WriteRune(r)
	}

	result := strings.ToLower(finalRes.String())

	// Clean up multiple underscores first
	for strings.Contains(result, "__") {
		result = strings.ReplaceAll(result, "__", "_")
	}

	// Then apply acronym replacements on the cleaned snake_case string
	for _, acronym := range acronyms {
		lowerAcronym := strings.ToLower(acronym)
		// Replace the acronym with underscores (e.g., "c_s_m" -> "csm")
		parts := strings.Split(lowerAcronym, "")
		withUnderscores := strings.Join(parts, "_")
		result = strings.ReplaceAll(result, withUnderscores, lowerAcronym)
	}

	// Trim any leading/trailing underscores
	result = strings.Trim(result, "_")
	return result
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

	return os.WriteFile("cmd/root.gen.go", formatted, 0644)
}
