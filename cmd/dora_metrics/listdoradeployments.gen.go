package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDORADeploymentsCmd = &cobra.Command{
	Use:   "listdoradeployments",
	Short: "Get a list of deployment events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dora/deployments")
		fmt.Println("OperationID: ListDORADeployments")
	},
}

func init() {
	Cmd.AddCommand(ListDORADeploymentsCmd)
}
