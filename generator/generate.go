package main

//go:generate go run .

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if err := RunGenerate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func RunGenerate() error {
	config, err := loadConfig("generator/config.yaml")
	if err != nil {
		return err
	}

	if err := cleanupGeneratedFiles(); err != nil {
		return err
	}

	spec, err := downloadOpenAPI()
	if err != nil {
		return err
	}

	tmpl, entrypointTmpl, err := parseTemplates(config)
	if err != nil {
		return err
	}

	tags := make(map[string]bool)
	skipOpsMap := make(map[string]bool)
	for _, op := range config.SkipOperations {
		skipOpsMap[op] = true
	}

	for path, methods := range spec.Paths {
		for method, op := range methods {
			if op.OperationID == "" || skipOpsMap[op.OperationID] {
				continue
			}

			tag, rawTag := normalizeTag(op)
			tags[tag] = true
			apiTagName := normalizeApiTagName(rawTag)

			pkgDir := filepath.Join("cmd", tag)
			if err := os.MkdirAll(pkgDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", pkgDir, err)
			}

			if err := ensureEntrypoint(pkgDir, tag, entrypointTmpl, config); err != nil {
				return err
			}

			data := prepareTemplateData(tag, rawTag, apiTagName, method, path, op, spec, config)

			if err := generateCommandFile(pkgDir, op.OperationID, data, tmpl, config); err != nil {
				return err
			}
		}
	}

	return updateRootGo(tags)
}
