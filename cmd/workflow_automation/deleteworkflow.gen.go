package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteWorkflowCmd = &cobra.Command{
	Use:   "deleteworkflow",
	Short: "Delete an existing Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/workflows/{workflow_id}")
		fmt.Println("OperationID: DeleteWorkflow")
	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowCmd)
}
