package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDeploymentGateCmd = &cobra.Command{
	Use:   "updatedeploymentgate",
	Short: "Update deployment gate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/deployment_gates/{id}")
		fmt.Println("OperationID: UpdateDeploymentGate")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeploymentGateCmd)
}
