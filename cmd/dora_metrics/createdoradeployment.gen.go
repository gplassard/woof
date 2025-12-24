package dora_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDORADeploymentCmd = &cobra.Command{
	Use:   "createdoradeployment",
	Short: "Send a deployment event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/dora/deployment")
		fmt.Println("OperationID: CreateDORADeployment")
	},
}

func init() {
	Cmd.AddCommand(CreateDORADeploymentCmd)
}
