package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	if err := runGenerate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

type OpenAPI struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

type Operation struct {
	OperationID string   `yaml:"operationId"`
	Summary     string   `yaml:"summary"`
	Tags        []string `yaml:"tags"`
}

type TemplateData struct {
	PackageName string
	CommandName string
	Use         string
	Short       string
	Method      string
	Path        string
	OperationID string
}

func runGenerate() error {
	// Cleanup existing generated files
	os.Remove("cmd/root.gen.go")
	entries, err := os.ReadDir("cmd")
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				pkgDir := filepath.Join("cmd", entry.Name())
				files, err := os.ReadDir(pkgDir)
				if err == nil {
					for _, f := range files {
						if strings.HasSuffix(f.Name(), ".gen.go") {
							os.Remove(filepath.Join(pkgDir, f.Name()))
						}
					}
				}
			}
		}
	}

	url := "https://raw.githubusercontent.com/DataDog/datadog-api-client-go/master/.generator/schemas/v2/openapi.yaml"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download openapi spec: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var spec OpenAPI
	if err := yaml.Unmarshal(body, &spec); err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	tmpl, err := template.ParseFiles("generator/command.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse command template: %w", err)
	}

	entrypointTmpl, err := template.ParseFiles("generator/entrypoint.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse entrypoint template: %w", err)
	}

	tags := make(map[string]bool)

	for path, methods := range spec.Paths {
		for method, op := range methods {
			if op.OperationID == "" {
				continue
			}

			tag := "general"
			if len(op.Tags) > 0 {
				tag = strings.ToLower(op.Tags[0])
				tag = strings.ReplaceAll(tag, " ", "_")
				tag = strings.ReplaceAll(tag, "-", "_")
			}
			tags[tag] = true

			pkgDir := filepath.Join("cmd", tag)
			if err := os.MkdirAll(pkgDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", pkgDir, err)
			}

			// Ensure entrypoint.gen.go exists in the package
			entrypointPath := filepath.Join(pkgDir, "entrypoint.gen.go")
			if _, err := os.Stat(entrypointPath); os.IsNotExist(err) {
				f, err := os.Create(entrypointPath)
				if err != nil {
					return fmt.Errorf("failed to create entrypoint for %s: %w", tag, err)
				}
				data := TemplateData{
					PackageName: tag,
					Use:         tag,
					Short:       fmt.Sprintf("%s endpoints", tag),
				}
				if err := entrypointTmpl.Execute(f, data); err != nil {
					f.Close()
					return fmt.Errorf("failed to execute entrypoint template for %s: %w", tag, err)
				}
				f.Close()
			}

			fileName := fmt.Sprintf("%s.gen.go", strings.ToLower(op.OperationID))
			filePath := filepath.Join(pkgDir, fileName)

			data := TemplateData{
				PackageName: tag,
				CommandName: strings.ToUpper(op.OperationID[:1]) + op.OperationID[1:] + "Cmd",
				Use:         strings.ToLower(op.OperationID),
				Short:       op.Summary,
				Method:      strings.ToUpper(method),
				Path:        path,
				OperationID: op.OperationID,
			}

			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", filePath, err)
			}

			if err := tmpl.Execute(f, data); err != nil {
				f.Close()
				return fmt.Errorf("failed to execute template for %s: %w", filePath, err)
			}
			f.Close()
			fmt.Printf("Generated %s\n", filePath)
		}
	}

	return updateRootGo(tags)
}

type RootTemplateData struct {
	Imports  []string
	Commands []string
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
