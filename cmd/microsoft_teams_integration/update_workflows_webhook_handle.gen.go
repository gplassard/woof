package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use: "update-workflows-webhook-handle [handle_id] [payload]",

	Short: "Update Workflows webhook handle",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.MicrosoftTeamsUpdateWorkflowsWebhookHandleRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.UpdateWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-workflows-webhook-handle")

		cmd.Println(cmdutil.FormatJSON(res, "workflows-webhook-handle"))
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowsWebhookHandleCmd)
}
