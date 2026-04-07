package main

//go:generate go run .

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if err := RunGenerate(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func RunGenerate(args []string) error {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}
	// Ensure we are working from the project root if we are in the generator directory
	if cwd, err := os.Getwd(); err == nil && filepath.Base(cwd) == "generator" {
		if err := os.Chdir(".."); err != nil {
			return fmt.Errorf("failed to change directory to project root: %w", err)
		}
		defer func() {
			_ = os.Chdir(cwd)
		}()
	}

	config, err := loadConfig()
	if err != nil {
		return err
	}

	if filter == "" {
		if err := cleanupGeneratedFiles(); err != nil {
			return err
		}
	}

	input, err := loadFromSDK(config)
	if err != nil {
		return err
	}

	tmpl, entrypointTmpl, err := parseTemplates(config)
	if err != nil {
		return err
	}

	bundles := make(map[string]bool)
	skipOpsMap := make(map[string]bool)
	for _, op := range config.SkipOperations {
		skipOpsMap[op] = true
	}

	for _, op := range input.Operations {
		if op.OperationID == "" || skipOpsMap[op.OperationID] {
			continue
		}

		if filter != "" && filter != op.OperationID {
			continue
		}

		bundles[op.Bundle] = true
		pkgDir := filepath.Join("cmd", op.Bundle)
		if err := os.MkdirAll(pkgDir, 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", pkgDir, err)
		}

		if err := ensureEntrypoint(pkgDir, op.Bundle, entrypointTmpl, config); err != nil {
			return err
		}

		data := prepareTemplateData(op, config)
		if err := generateCommandFile(pkgDir, op.OperationID, data, tmpl, config); err != nil {
			return err
		}
	}

	if filter == "" {
		return updateRootGo(bundles)
	}

	return nil
}
