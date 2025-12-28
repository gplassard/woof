package opsgenie_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOpsgenieServiceCmd = &cobra.Command{
	Use:   "get-opsgenie-service [integration_service_id]",
	
	Short: "Get a single service object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetOpsgenieService(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-opsgenie-service: %v", err)
		}

		cmdutil.PrintJSON(res, "opsgenie-service")
	},
}

func init() {
	Cmd.AddCommand(GetOpsgenieServiceCmd)
}
