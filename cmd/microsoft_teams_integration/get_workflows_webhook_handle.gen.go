package microsoft_teams_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "get-workflows-webhook-handle [handle_id]",
	
	Short: "Get Workflows webhook handle information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-workflows-webhook-handle: %v", err)
		}

		cmdutil.PrintJSON(res, "workflows-webhook-handle")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowsWebhookHandleCmd)
}
