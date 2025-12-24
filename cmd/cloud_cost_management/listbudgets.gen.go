package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListBudgetsCmd = &cobra.Command{
	Use:   "listbudgets",
	Short: "List budgets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/budgets")
		fmt.Println("OperationID: ListBudgets")
	},
}

func init() {
	Cmd.AddCommand(ListBudgetsCmd)
}
