package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDeploymentGateCmd = &cobra.Command{
	Use:   "getdeploymentgate",
	Short: "Get deployment gate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/deployment_gates/{id}")
		fmt.Println("OperationID: GetDeploymentGate")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentGateCmd)
}
