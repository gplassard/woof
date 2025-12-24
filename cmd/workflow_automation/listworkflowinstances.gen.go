package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListWorkflowInstancesCmd = &cobra.Command{
	Use:   "listworkflowinstances",
	Short: "List workflow instances",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/workflows/{workflow_id}/instances")
		fmt.Println("OperationID: ListWorkflowInstances")
	},
}

func init() {
	Cmd.AddCommand(ListWorkflowInstancesCmd)
}
