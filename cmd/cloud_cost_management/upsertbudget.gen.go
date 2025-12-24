package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpsertBudgetCmd = &cobra.Command{
	Use:   "upsertbudget",
	Short: "Create or update a budget",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/cost/budget")
		fmt.Println("OperationID: UpsertBudget")
	},
}

func init() {
	Cmd.AddCommand(UpsertBudgetCmd)
}
