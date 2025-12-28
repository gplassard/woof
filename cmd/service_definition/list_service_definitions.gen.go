package service_definition

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListServiceDefinitionsCmd = &cobra.Command{
	Use:   "list-service-definitions",
	Aliases: []string{ "list-s", },
	Short: "Get all service definitions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err := api.ListServiceDefinitions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-service-definitions")

		cmdutil.PrintJSON(res, "service_definition")
	},
}

func init() {
	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
