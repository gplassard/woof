package microsoft_teams_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use: "update-workflows-webhook-handle [handle_id]",

	Short: "Update Workflows webhook handle",
	Long: `Update Workflows webhook handle
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#update-workflows-webhook-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsWorkflowsWebhookHandleResponse
		var err error

		var body datadogV2.MicrosoftTeamsUpdateWorkflowsWebhookHandleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-workflows-webhook-handle")

		fmt.Println(cmdutil.FormatJSON(res, "workflows_webhook_handle"))
	},
}

func init() {

	UpdateWorkflowsWebhookHandleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateWorkflowsWebhookHandleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateWorkflowsWebhookHandleCmd)
}
