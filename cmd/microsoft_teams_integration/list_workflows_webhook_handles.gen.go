package microsoft_teams_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListWorkflowsWebhookHandlesCmd = &cobra.Command{
	Use: "list-workflows-webhook-handles",

	Short: "Get all Workflows webhook handles",
	Long: `Get all Workflows webhook handles
Documentation: https://docs.datadoghq.com/api/latest/microsoft-teams-integration/#list-workflows-webhook-handles`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MicrosoftTeamsWorkflowsWebhookHandlesResponse
		var err error

		optionalParams := datadogV2.NewListWorkflowsWebhookHandlesOptionalParameters()

		if cmd.Flags().Changed("name") {
			val, _ := cmd.Flags().GetString("name")
			optionalParams.WithName(val)
		}

		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListWorkflowsWebhookHandles(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-workflows-webhook-handles")

		fmt.Println(cmdutil.FormatJSON(res, "workflows-webhook-handle"))
	},
}

func init() {

	ListWorkflowsWebhookHandlesCmd.Flags().String("name", "", "Your Workflows webhook handle name.")

	Cmd.AddCommand(ListWorkflowsWebhookHandlesCmd)
}
