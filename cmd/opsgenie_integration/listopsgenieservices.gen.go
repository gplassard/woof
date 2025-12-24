package opsgenie_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListOpsgenieServicesCmd = &cobra.Command{
	Use:   "listopsgenieservices",
	Short: "Get all service objects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/opsgenie/services")
		fmt.Println("OperationID: ListOpsgenieServices")
	},
}

func init() {
	Cmd.AddCommand(ListOpsgenieServicesCmd)
}
