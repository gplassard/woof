package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDeploymentGateCmd = &cobra.Command{
	Use:   "createdeploymentgate",
	Short: "Create deployment gate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/deployment_gates")
		fmt.Println("OperationID: CreateDeploymentGate")
	},
}

func init() {
	Cmd.AddCommand(CreateDeploymentGateCmd)
}
