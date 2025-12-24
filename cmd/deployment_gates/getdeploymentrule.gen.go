package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDeploymentRuleCmd = &cobra.Command{
	Use:   "getdeploymentrule",
	Short: "Get deployment rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/deployment_gates/{gate_id}/rules/{id}")
		fmt.Println("OperationID: GetDeploymentRule")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentRuleCmd)
}
