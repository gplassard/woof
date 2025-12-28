package opsgenie_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOpsgenieServiceCmd = &cobra.Command{
	Use:   "delete-opsgenie-service [integration_service_id]",
	Short: "Delete a single service object",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteOpsgenieService(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-opsgenie-service: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOpsgenieServiceCmd)
}
