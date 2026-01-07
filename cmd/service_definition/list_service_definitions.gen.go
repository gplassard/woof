package service_definition

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListServiceDefinitionsCmd = &cobra.Command{
	Use:     "list-service-definitions",
	Aliases: []string{"list"},
	Short:   "Get all service definitions",
	Long: `Get all service definitions
Documentation: https://docs.datadoghq.com/api/latest/service-definition/#list-service-definitions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceDefinitionsListResponse
		var err error

		optionalParams := datadogV2.NewListServiceDefinitionsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("schema-version") {
			val, _ := cmd.Flags().GetString("schema-version")
			optionalParams.WithSchemaVersion(val)
		}

		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceDefinitions(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-service-definitions")

		cmd.Println(cmdutil.FormatJSON(res, "service_definition"))
	},
}

func init() {

	ListServiceDefinitionsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListServiceDefinitionsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListServiceDefinitionsCmd.Flags().String("schema-version", "", "The schema version desired in the response.")

	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
