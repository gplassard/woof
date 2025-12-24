package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDeploymentRuleCmd = &cobra.Command{
	Use:   "updatedeploymentrule",
	Short: "Update deployment rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/deployment_gates/{gate_id}/rules/{id}")
		fmt.Println("OperationID: UpdateDeploymentRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeploymentRuleCmd)
}
