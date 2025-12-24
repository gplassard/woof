package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDeploymentGateCmd = &cobra.Command{
	Use:   "deletedeploymentgate",
	Short: "Delete deployment gate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/deployment_gates/{id}")
		fmt.Println("OperationID: DeleteDeploymentGate")
	},
}

func init() {
	Cmd.AddCommand(DeleteDeploymentGateCmd)
}
