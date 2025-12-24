package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetProjectedCostCmd = &cobra.Command{
	Use:   "getprojectedcost",
	Short: "Get projected cost across your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/projected_cost")
		fmt.Println("OperationID: GetProjectedCost")
	},
}

func init() {
	Cmd.AddCommand(GetProjectedCostCmd)
}
