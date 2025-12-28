package opsgenie_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOpsgenieServicesCmd = &cobra.Command{
	Use:   "list-opsgenie-services",
	
	Short: "Get all service objects",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListOpsgenieServices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-opsgenie-services")

		cmdutil.PrintJSON(res, "opsgenie-service")
	},
}

func init() {
	Cmd.AddCommand(ListOpsgenieServicesCmd)
}
