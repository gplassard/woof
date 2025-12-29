package microsoft_teams_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "update-workflows-webhook-handle [handle_id]",
	
	Short: "Update Workflows webhook handle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MicrosoftTeamsUpdateWorkflowsWebhookHandleRequest{})
		cmdutil.HandleError(err, "failed to update-workflows-webhook-handle")

		cmd.Println(cmdutil.FormatJSON(res, "workflows-webhook-handle"))
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowsWebhookHandleCmd)
}
