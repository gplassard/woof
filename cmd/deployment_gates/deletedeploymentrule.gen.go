package deployment_gates

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDeploymentRuleCmd = &cobra.Command{
	Use:   "deletedeploymentrule",
	Short: "Delete deployment rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/deployment_gates/{gate_id}/rules/{id}")
		fmt.Println("OperationID: DeleteDeploymentRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteDeploymentRuleCmd)
}
