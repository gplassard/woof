package microsoft_teams_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "getworkflowswebhookhandle [handle_id]",
	Short: "Get Workflows webhook handle information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMicrosoftTeamsIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetWorkflowsWebhookHandle(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getworkflowswebhookhandle: %v", err)
		}

		cmdutil.PrintJSON(res, "microsoft_teams_integration")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowsWebhookHandleCmd)
}
