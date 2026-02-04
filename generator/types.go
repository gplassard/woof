package main

type Config struct {
	OnlyBundles                  []string            `yaml:"only_bundles"`
	SkipOperations               []string            `yaml:"skip_operations"`
	OptionalParametersOperations map[string]string   `yaml:"optional_parameters_operations"`
	BundleAliases                map[string][]string `yaml:"bundle_aliases"`
	Acronyms                     []string            `yaml:"acronyms"`
}

// SDK-based types for parsing the Go SDK

type SDKSpec struct {
	Operations map[string]*SDKOperation
}

type SDKOperation struct {
	OperationID       string
	Summary           string
	Description       string
	Tags              []string
	Method            string // GET, POST, PUT, PATCH, DELETE
	Path              string
	APIName           string // e.g., "action_connection"
	Parameters        []SDKParameter
	HasOptionalParams bool
	RequestBody       *SDKRequestBody
	ResponseType      string
	HasResponse       bool
}

type SDKParameter struct {
	Name     string
	Type     string
	In       string // "path", "query", "body"
	Required bool
}

type SDKRequestBody struct {
	Required bool
	Type     string
}

// Legacy OpenAPI types (kept for backward compatibility during migration)

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
			Schema map[string]interface{} `yaml:",inline"`
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
				Ref string `yaml:"$ref"`
			} `yaml:"schema"`
		} `yaml:"content"`
	} `yaml:"requestBody"`
	Responses map[string]struct {
		Content map[string]interface{} `yaml:"content"`
	} `yaml:"responses"`
}

type TemplateData struct {
	PackageName      string
	CommandName      string
	Use              string
	Short            string
	Method           string
	Path             string
	OperationID      string
	BundleName       string
	ApiBundleName    string
	Args             []string
	ArgTypes         []string
	HasRequestBody   bool
	RequestBodyType  string
	IsOptionalParams bool
	HasResponse      bool
	ResourceType     string
	ResponseTypeGo   string
	Aliases          []string
	DocURL           string
}

type RootTemplateData struct {
	Imports  []string
	Commands []string
}
