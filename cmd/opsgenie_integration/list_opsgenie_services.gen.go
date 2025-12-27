package opsgenie_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOpsgenieServicesCmd = &cobra.Command{
	Use:   "list_opsgenie_services",
	Short: "Get all service objects",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListOpsgenieServices(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_opsgenie_services: %v", err)
		}

		cmdutil.PrintJSON(res, "opsgenie_integration")
	},
}

func init() {
	Cmd.AddCommand(ListOpsgenieServicesCmd)
}
