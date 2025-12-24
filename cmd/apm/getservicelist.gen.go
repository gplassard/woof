package apm

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetServiceListCmd = &cobra.Command{
	Use:   "getservicelist",
	Short: "Get service list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apm/services")
		fmt.Println("OperationID: GetServiceList")
	},
}

func init() {
	Cmd.AddCommand(GetServiceListCmd)
}
