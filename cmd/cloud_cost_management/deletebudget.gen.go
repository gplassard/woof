package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteBudgetCmd = &cobra.Command{
	Use:   "deletebudget",
	Short: "Delete a budget",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/budget/{budget_id}")
		fmt.Println("OperationID: DeleteBudget")
	},
}

func init() {
	Cmd.AddCommand(DeleteBudgetCmd)
}
