package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "create-workflows-webhook-handle",
	
	Short: "Create Workflows webhook handle",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), datadogV2.MicrosoftTeamsCreateWorkflowsWebhookHandleRequest{})
		if err != nil {
			log.Fatalf("failed to create-workflows-webhook-handle: %v", err)
		}

		cmdutil.PrintJSON(res, "workflows-webhook-handle")
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowsWebhookHandleCmd)
}
