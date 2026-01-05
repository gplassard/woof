package microsoft_teams_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListWorkflowsWebhookHandlesCmd = &cobra.Command{
	Use: "list-workflows-webhook-handles",

	Short: "Get all Workflows webhook handles",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListWorkflowsWebhookHandles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-workflows-webhook-handles")

		cmd.Println(cmdutil.FormatJSON(res, "workflows-webhook-handle"))
	},
}

func init() {
	Cmd.AddCommand(ListWorkflowsWebhookHandlesCmd)
}
