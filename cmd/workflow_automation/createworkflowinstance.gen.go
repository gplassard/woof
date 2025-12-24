package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateWorkflowInstanceCmd = &cobra.Command{
	Use:   "createworkflowinstance",
	Short: "Execute a workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/workflows/{workflow_id}/instances")
		fmt.Println("OperationID: CreateWorkflowInstance")
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowInstanceCmd)
}
