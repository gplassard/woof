package service_definition

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListServiceDefinitionsCmd = &cobra.Command{
	Use:   "list_service_definitions",
	Short: "Get all service definitions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err := api.ListServiceDefinitions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_service_definitions: %v", err)
		}

		cmdutil.PrintJSON(res, "service_definition")
	},
}

func init() {
	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
