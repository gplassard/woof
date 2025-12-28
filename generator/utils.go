package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

func loadConfig(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &config, nil
}

func cleanupGeneratedFiles() error {
	os.Remove("cmd/root.gen.go")
	entries, err := os.ReadDir("cmd")
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			pkgDir := filepath.Join("cmd", entry.Name())
			os.RemoveAll(pkgDir)
		}
	}
	return nil
}

func downloadOpenAPI() (*OpenAPI, error) {
	url := "https://raw.githubusercontent.com/DataDog/datadog-api-client-go/master/.generator/schemas/v2/openapi.yaml"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download openapi spec: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var spec OpenAPI
	if err := yaml.Unmarshal(body, &spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}
	return &spec, nil
}

func parseTemplates(config *Config) (*template.Template, *template.Template, error) {
	funcMap := template.FuncMap{
		"replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"lower": strings.ToLower,
		"toSnakeCase": func(s string) string {
			return toSnakeCase(s, config)
		},
		"toKebabCase": func(s string) string {
			return toKebabCase(s, config)
		},
	}

	tmpl, err := template.New("command.go.tmpl").Funcs(funcMap).ParseFiles("generator/command.go.tmpl")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse command template: %w", err)
	}

	entrypointTmpl, err := template.New("entrypoint.go.tmpl").Funcs(funcMap).ParseFiles("generator/entrypoint.go.tmpl")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse entrypoint template: %w", err)
	}

	return tmpl, entrypointTmpl, nil
}

func normalizeTag(op Operation) (string, string) {
	rawTag := "general"
	if len(op.Tags) > 0 {
		rawTag = op.Tags[0]
	}
	tag := strings.ToLower(rawTag)
	tag = strings.ReplaceAll(tag, " ", "_")
	tag = strings.ReplaceAll(tag, "-", "_")
	return tag, rawTag
}

func normalizeApiTagName(rawTag string) string {
	name := strings.ReplaceAll(rawTag, " ", "")
	return strings.ReplaceAll(name, "-", "")
}
