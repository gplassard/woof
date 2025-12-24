package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetHistoricalCostByOrgCmd = &cobra.Command{
	Use:   "gethistoricalcostbyorg",
	Short: "Get historical cost across your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/historical_cost")
		fmt.Println("OperationID: GetHistoricalCostByOrg")
	},
}

func init() {
	Cmd.AddCommand(GetHistoricalCostByOrgCmd)
}
