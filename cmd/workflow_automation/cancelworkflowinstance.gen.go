package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CancelWorkflowInstanceCmd = &cobra.Command{
	Use:   "cancelworkflowinstance",
	Short: "Cancel a workflow instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/workflows/{workflow_id}/instances/{instance_id}/cancel")
		fmt.Println("OperationID: CancelWorkflowInstance")
	},
}

func init() {
	Cmd.AddCommand(CancelWorkflowInstanceCmd)
}
