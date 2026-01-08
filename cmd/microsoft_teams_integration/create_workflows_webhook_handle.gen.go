package microsoft_teams_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use: "create-workflows-webhook-handle",

	Short: "Create Workflows webhook handle",
	Long: `Create Workflows webhook handle
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#create-workflows-webhook-handle`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsWorkflowsWebhookHandleResponse
		var err error

		var body datadogV2.MicrosoftTeamsCreateWorkflowsWebhookHandleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-workflows-webhook-handle")

		fmt.Println(cmdutil.FormatJSON(res, "workflows_webhook_handle"))
	},
}

func init() {

	CreateWorkflowsWebhookHandleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateWorkflowsWebhookHandleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateWorkflowsWebhookHandleCmd)
}
