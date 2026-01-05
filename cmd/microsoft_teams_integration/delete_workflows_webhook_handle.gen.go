package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteWorkflowsWebhookHandleCmd = &cobra.Command{
	Use: "delete-workflows-webhook-handle [handle_id]",

	Short: "Delete Workflows webhook handle",
	Long: `Delete Workflows webhook handle
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#delete-workflows-webhook-handle`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		_, err = api.DeleteWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-workflows-webhook-handle")

	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowsWebhookHandleCmd)
}
