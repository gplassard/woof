package service_definition

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListServiceDefinitionsCmd = &cobra.Command{
	Use:   "listservicedefinitions",
	Short: "Get all service definitions",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err := api.ListServiceDefinitions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listservicedefinitions: %v", err)
		}

		cmdutil.PrintJSON(res, "service_definition")
	},
}

func init() {
	Cmd.AddCommand(ListServiceDefinitionsCmd)
}
