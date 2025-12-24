package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "getworkflowswebhookhandle",
	Short: "Get Workflows webhook handle information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/ms-teams/configuration/workflows-webhook-handles/{handle_id}")
		fmt.Println("OperationID: GetWorkflowsWebhookHandle")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowsWebhookHandleCmd)
}
