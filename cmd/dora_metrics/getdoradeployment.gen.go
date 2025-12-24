package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDORADeploymentCmd = &cobra.Command{
	Use:   "getdoradeployment",
	Short: "Get a deployment event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/dora/deployments/{deployment_id}")
		fmt.Println("OperationID: GetDORADeployment")
	},
}

func init() {
	Cmd.AddCommand(GetDORADeploymentCmd)
}
