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
	Paths      map[string]map[string]Operation `yaml:"paths"`
	Components struct {
		Parameters map[string]struct {
			Name     string                 `yaml:"name"`
			In       string                 `yaml:"in"`
			Required bool                   `yaml:"required"`
			Schema   map[string]interface{} `yaml:"schema"`
		} `yaml:"parameters"`
		Schemas map[string]struct {
			Type   string                 `yaml:"type"`
			Format string                 `yaml:"format"`
			Ref    string                 `yaml:"$ref"`
			Schema map[string]interface{} `yaml:"schema"`
		} `yaml:"schemas"`
	} `yaml:"components"`
}

type Operation struct {
	OperationID string   `yaml:"operationId"`
	Summary     string   `yaml:"summary"`
	Tags        []string `yaml:"tags"`
	Parameters  []struct {
		Name     string                 `yaml:"name"`
		In       string                 `yaml:"in"`
		Required bool                   `yaml:"required"`
		Ref      string                 `yaml:"$ref"`
		Schema   map[string]interface{} `yaml:"schema"`
	} `yaml:"parameters"`
	RequestBody *struct {
		Required bool `yaml:"required"`
		Content  map[string]struct {
			Schema struct {
				Ref        string `yaml:"$ref"`
				Extensions struct {
					GoName string `yaml:"x-go-name"`
				} `yaml:"x-go-name"`
			} `yaml:"schema"`
		} `yaml:"content"`
		Extensions struct {
			GoName string `yaml:"x-go-name"`
		} `yaml:"x-go-name"`
	} `yaml:"requestBody"`
	Responses map[string]struct {
		Content map[string]interface{} `yaml:"content"`
	} `yaml:"responses"`
}

type TemplateData struct {
	PackageName        string
	CommandName        string
	Use                string
	Short              string
	Method             string
	Path               string
	OperationID        string
	TagName            string
	ApiTagName         string
	Args               []string
	ArgTypes           []string
	HasRequestBody     bool
	RequestBodyType    string
	IsOptionalParams   bool
	HasResponse        bool
	ReturnsThreeValues bool
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

	funcMap := template.FuncMap{
		"replace": func(old, new, src string) string {
			return strings.ReplaceAll(src, old, new)
		},
		"lower": strings.ToLower,
	}

	tmpl, err := template.New("command.go.tmpl").Funcs(funcMap).ParseFiles("generator/command.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse command template: %w", err)
	}

	entrypointTmpl, err := template.New("entrypoint.go.tmpl").Funcs(funcMap).ParseFiles("generator/entrypoint.go.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse entrypoint template: %w", err)
	}

	tags := make(map[string]bool)

	for path, methods := range spec.Paths {
		for method, op := range methods {
			// Hardcoded list of operations to skip (e.g., missing from SDK)
			skipOps := map[string]bool{
				"CreateIncidentAttachment":   true,
				"DeleteIncidentAttachment":   true,
				"UpdateIncidentAttachment":   true,
				"CreateUserNotificationRule": true,
				"DeleteUserNotificationRule": true,
				"GetUserNotificationRule":    true,
				"ListUserNotificationRules":  true,
				"UpdateUserNotificationRule": true,
				"MakeGCPSTSDelegateRequest":  true, // This one might be a type mismatch but skipping for now
				"UploadCustomCostsFile":      true, // Undefined type
				"MakeGCPSTSDelegate":         true,
				"SubmitLog":                  true, // Undefined type HTTPLog
			}
			if skipOps[op.OperationID] {
				continue
			}

			tag := "general"
			rawTag := "general"
			if len(op.Tags) > 0 {
				rawTag = op.Tags[0]
				tag = strings.ToLower(rawTag)
				tag = strings.ReplaceAll(tag, " ", "_")
				tag = strings.ReplaceAll(tag, "-", "_")
			}
			tags[tag] = true

			apiTagName := strings.ReplaceAll(rawTag, " ", "")
			apiTagName = strings.ReplaceAll(apiTagName, "-", "")

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

			var args []string
			var argTypes []string
			use := strings.ToLower(op.OperationID)
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

					// Resolve schema if it's a ref
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
								iType, _ := items["type"].(string)
								if iType == "string" {
									argType = "[]string"
								}
							}
						}
					}

					// Custom type overrides for some known SDK arguments
					if op.OperationID == "GetSBOM" && name == "asset_type" {
						argType = "datadogV2.AssetType"
					}
					if op.OperationID == "GetTeamSync" && name == "filter[source]" {
						argType = "datadogV2.TeamSyncAttributesSource"
					}
					if op.OperationID == "GetSignalNotificationRules" || op.OperationID == "GetVulnerabilityNotificationRules" {
						// Special case, will be handled by hasResponse later
					}
					if op.OperationID == "GetCostByOrg" && name == "start_month" {
						argType = "time.Time"
					}
					if op.OperationID == "GetHistoricalCostByOrg" && name == "start_month" {
						argType = "time.Time"
					}
					if op.OperationID == "GetHourlyUsage" && name == "start_hr" {
						argType = "time.Time"
					}
					if op.OperationID == "GetMonthlyCostAttribution" && name == "start_month" {
						argType = "time.Time"
					}
					if op.OperationID == "GetUsageApplicationSecurityMonitoring" && name == "start_hr" {
						argType = "time.Time"
					}
					if op.OperationID == "GetUsageLambdaTracedInvocations" && name == "start_hr" {
						argType = "time.Time"
					}
					if op.OperationID == "GetUsageObservabilityPipelines" && name == "start_hr" {
						argType = "time.Time"
					}

					argTypes = append(argTypes, argType)
				}
			}

			fileName := fmt.Sprintf("%s.gen.go", strings.ToLower(op.OperationID))
			filePath := filepath.Join(pkgDir, fileName)

			hasResponse := false
			var responseCodes []string
			for code, resp := range op.Responses {
				responseCodes = append(responseCodes, code)
				if code == "200" || code == "201" || code == "202" {
					if resp.Content != nil {
						hasResponse = true
					}
				}
			}
			sort.Strings(responseCodes)

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

				// Hardcoded list of operations known to use OptionalParameters in the Go SDK
				// This is a workaround because the OpenAPI spec doesn't explicitly flag these
				optionalParamsOps := map[string]string{
					"SearchCIAppTestEvents":               "SearchCIAppTestEventsOptionalParameters",
					"SearchCIAppPipelineEvents":           "SearchCIAppPipelineEventsOptionalParameters",
					"SearchAuditLogs":                     "SearchAuditLogsOptionalParameters",
					"SearchEvents":                        "SearchEventsOptionalParameters",
					"CreateOpenAPI":                       "CreateOpenAPIOptionalParameters",
					"UpdateOpenAPI":                       "UpdateOpenAPIOptionalParameters",
					"ListLogs":                            "ListLogsOptionalParameters",
					"UploadIdPMetadata":                   "UploadIdPMetadataOptionalParameters",
					"SearchFlakyTests":                    "SearchFlakyTestsOptionalParameters",
					"SearchSecurityMonitoringHistsignals": "SearchSecurityMonitoringHistsignalsOptionalParameters",
					"SearchSecurityMonitoringSignals":     "SearchSecurityMonitoringSignalsOptionalParameters",
				}

				if optType, ok := optionalParamsOps[op.OperationID]; ok {
					isOptionalParams = true
					requestBodyType = optType
				}
			}

			data := TemplateData{
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
