package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDeploymentRuleCmd = &cobra.Command{
	Use:   "createdeploymentrule",
	Short: "Create deployment rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/deployment_gates/{gate_id}/rules")
		fmt.Println("OperationID: CreateDeploymentRule")
	},
}

func init() {
	Cmd.AddCommand(CreateDeploymentRuleCmd)
}
