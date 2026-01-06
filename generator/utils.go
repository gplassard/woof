package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
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

func downloadOpenAPI() (*OpenAPI, error) {
	versionContent, err := generatorFS.ReadFile("version.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read version file: %w", err)
	}
	var version Version
	if err := yaml.Unmarshal(versionContent, &version); err != nil {
		return nil, fmt.Errorf("failed to unmarshal version: %w", err)
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/DataDog/datadog-api-client-go/%s/.generator/schemas/v2/openapi.yaml", version.Sha)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download openapi spec: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

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

func normalizeBundle(op Operation) (string, string) {
	rawBundle := "general"
	if len(op.Tags) > 0 {
		rawBundle = op.Tags[0]
	}
	bundle := strings.ToLower(rawBundle)
	bundle = strings.ReplaceAll(bundle, " ", "_")
	bundle = strings.ReplaceAll(bundle, "-", "_")
	return bundle, rawBundle
}

func normalizeApiBundleName(rawBundle string) string {
	name := strings.ReplaceAll(rawBundle, " ", "")
	return strings.ReplaceAll(name, "-", "")
}
