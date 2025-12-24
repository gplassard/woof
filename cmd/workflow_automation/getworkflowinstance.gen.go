package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetWorkflowInstanceCmd = &cobra.Command{
	Use:   "getworkflowinstance",
	Short: "Get a workflow instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/workflows/{workflow_id}/instances/{instance_id}")
		fmt.Println("OperationID: GetWorkflowInstance")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowInstanceCmd)
}
