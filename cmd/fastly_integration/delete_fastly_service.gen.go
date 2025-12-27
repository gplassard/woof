package fastly_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteFastlyServiceCmd = &cobra.Command{
	Use:   "delete_fastly_service [account_id] [service_id]",
	Short: "Delete Fastly service",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteFastlyService(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete_fastly_service: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteFastlyServiceCmd)
}
