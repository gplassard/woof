package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDeploymentGateRulesCmd = &cobra.Command{
	Use:   "getdeploymentgaterules",
	Short: "Get rules for a deployment gate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/deployment_gates/{gate_id}/rules")
		fmt.Println("OperationID: GetDeploymentGateRules")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentGateRulesCmd)
}
