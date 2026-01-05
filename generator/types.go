package main

type Config struct {
	SkipOperations               []string            `yaml:"skip_operations"`
	OptionalParametersOperations map[string]string   `yaml:"optional_parameters_operations"`
	BundleAliases                map[string][]string `yaml:"bundle_aliases"`
	Acronyms                     []string            `yaml:"acronyms"`
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
