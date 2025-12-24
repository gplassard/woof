package opsgenie_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOpsgenieServiceCmd = &cobra.Command{
	Use:   "createopsgenieservice",
	Short: "Create a new service object",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/opsgenie/services")
		fmt.Println("OperationID: CreateOpsgenieService")
	},
}

func init() {
	Cmd.AddCommand(CreateOpsgenieServiceCmd)
}
