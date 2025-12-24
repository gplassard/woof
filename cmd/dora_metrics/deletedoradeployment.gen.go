package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDORADeploymentCmd = &cobra.Command{
	Use:   "deletedoradeployment",
	Short: "Delete a deployment event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/dora/deployment/{deployment_id}")
		fmt.Println("OperationID: DeleteDORADeployment")
	},
}

func init() {
	Cmd.AddCommand(DeleteDORADeploymentCmd)
}
