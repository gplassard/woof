package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateWorkflowCmd = &cobra.Command{
	Use:   "updateworkflow",
	Short: "Update an existing Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/workflows/{workflow_id}")
		fmt.Println("OperationID: UpdateWorkflow")
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowCmd)
}
