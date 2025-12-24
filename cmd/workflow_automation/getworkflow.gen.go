package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetWorkflowCmd = &cobra.Command{
	Use:   "getworkflow",
	Short: "Get an existing Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/workflows/{workflow_id}")
		fmt.Println("OperationID: GetWorkflow")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowCmd)
}
