package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetEstimatedCostByOrgCmd = &cobra.Command{
	Use:   "getestimatedcostbyorg",
	Short: "Get estimated cost across your account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/estimated_cost")
		fmt.Println("OperationID: GetEstimatedCostByOrg")
	},
}

func init() {
	Cmd.AddCommand(GetEstimatedCostByOrgCmd)
}
