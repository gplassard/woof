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

		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceDefinitions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-service-definitions")

		cmd.Println(cmdutil.FormatJSON(res, "service_definition"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
