package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "updateworkflowswebhookhandle",
	Short: "Update Workflows webhook handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integration/ms-teams/configuration/workflows-webhook-handles/{handle_id}")
		fmt.Println("OperationID: UpdateWorkflowsWebhookHandle")
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowsWebhookHandleCmd)
}
