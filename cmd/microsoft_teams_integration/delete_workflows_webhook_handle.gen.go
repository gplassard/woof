package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "delete-workflows-webhook-handle [handle_id]",
	
	Short: "Delete Workflows webhook handle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		_, err := api.DeleteWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-workflows-webhook-handle: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowsWebhookHandleCmd)
}
