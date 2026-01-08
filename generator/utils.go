package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml command.go.tmpl entrypoint.go.tmpl root.go.tmpl version.json
var generatorFS embed.FS

type Version struct {
	Sha string `json:"sha"`
}

func loadConfig() (*Config, error) {
	content, err := generatorFS.ReadFile("config.yaml")
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
	return os.RemoveAll("cmd")
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
		"sub": func(a, b int) int {
			return a - b
		},
		"ternary": func(cond bool, a, b int) int {
			if cond {
				return a
			}
			return b
		},
		"len": func(v interface{}) int {
			if s, ok := v.([]string); ok {
				return len(s)
			}
			return 0
		},
		"contains": strings.Contains,
	}

	commandTmplContent, err := generatorFS.ReadFile("command.go.tmpl")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read command template: %w", err)
	}
	tmpl, err := template.New("command.go.tmpl").Funcs(funcMap).Parse(string(commandTmplContent))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse command template: %w", err)
	}

	entrypointTmplContent, err := generatorFS.ReadFile("entrypoint.go.tmpl")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read entrypoint template: %w", err)
	}
	entrypointTmpl, err := template.New("entrypoint.go.tmpl").Funcs(funcMap).Parse(string(entrypointTmplContent))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse entrypoint template: %w", err)
	}

	return tmpl, entrypointTmpl, nil
}

func normalizeApiBundleName(rawBundle string) string {
	name := strings.ReplaceAll(rawBundle, " ", "")
	return strings.ReplaceAll(name, "-", "")
}
