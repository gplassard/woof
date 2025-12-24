package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateWorkflowsWebhookHandleCmd = &cobra.Command{
	Use:   "createworkflowswebhookhandle",
	Short: "Create Workflows webhook handle",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/ms-teams/configuration/workflows-webhook-handles")
		fmt.Println("OperationID: CreateWorkflowsWebhookHandle")
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowsWebhookHandleCmd)
}
