package workflow_automation

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateWorkflowCmd = &cobra.Command{
	Use:   "createworkflow",
	Short: "Create a Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/workflows")
		fmt.Println("OperationID: CreateWorkflow")
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowCmd)
}
