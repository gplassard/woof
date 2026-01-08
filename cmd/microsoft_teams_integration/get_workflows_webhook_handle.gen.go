package microsoft_teams_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetWorkflowsWebhookHandleCmd = &cobra.Command{
	Use: "get-workflows-webhook-handle [handle_id]",

	Short: "Get Workflows webhook handle information",
	Long: `Get Workflows webhook handle information
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#get-workflows-webhook-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsWorkflowsWebhookHandleResponse
		var err error

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-workflows-webhook-handle")

		fmt.Println(cmdutil.FormatJSON(res, "workflows_webhook_handle"))
	},
}

func init() {

	Cmd.AddCommand(GetWorkflowsWebhookHandleCmd)
}
