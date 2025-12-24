package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetBudgetCmd = &cobra.Command{
	Use:   "getbudget",
	Short: "Get a budget",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/budget/{budget_id}")
		fmt.Println("OperationID: GetBudget")
	},
}

func init() {
	Cmd.AddCommand(GetBudgetCmd)
}
