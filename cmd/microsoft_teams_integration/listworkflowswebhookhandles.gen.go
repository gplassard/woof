package microsoft_teams_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListWorkflowsWebhookHandlesCmd = &cobra.Command{
	Use:   "listworkflowswebhookhandles",
	Short: "Get all Workflows webhook handles",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/ms-teams/configuration/workflows-webhook-handles")
		fmt.Println("OperationID: ListWorkflowsWebhookHandles")
	},
}

func init() {
	Cmd.AddCommand(ListWorkflowsWebhookHandlesCmd)
}
