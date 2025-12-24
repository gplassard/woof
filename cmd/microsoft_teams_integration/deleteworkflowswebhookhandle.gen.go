package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "deleteworkflowswebhookhandle",
	Short: "Delete Workflows webhook handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/ms-teams/configuration/workflows-webhook-handles/{handle_id}")
		fmt.Println("OperationID: DeleteWorkflowsWebhookHandle")
	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowsWebhookHandleCmd)
}
