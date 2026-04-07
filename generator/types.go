package main

type Config struct {
	SkipOperations               []string            `yaml:"skip_operations"`
	OptionalParametersOperations map[string]string   `yaml:"optional_parameters_operations"`
	BundleAliases                map[string][]string `yaml:"bundle_aliases"`
	Acronyms                     []string            `yaml:"acronyms"`
}

type GenerationInput struct {
	Operations []OperationModel
}

type OperationModel struct {
	OperationID      string
	Summary          string
	Method           string
	Path             string
	Bundle           string
	ApiBundleName    string
	Parameters       []ParameterModel
	HasRequestBody   bool
	RequestBodyType  string
	IsOptionalParams bool
	HasResponse      bool
	ResponseTypeGo   string
	ResourceType     string
}

type ParameterModel struct {
	GoName string
	Name   string
	GoType string
	In     string
}

type SDKModel struct {
	Services   []SDKService
	Operations []OperationModel
	Structs    map[string]SDKStruct
	Enums      map[string]SDKEnum
}

type SDKService struct {
	Name       string
	Bundle     string
	ApiName    string
	Operations []string
}

type SDKStruct struct {
	Name   string
	Fields map[string]SDKField
}

type SDKField struct {
	JSONName string
	TypeName string
}

type SDKEnum struct {
	Name   string
	Values []string
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
