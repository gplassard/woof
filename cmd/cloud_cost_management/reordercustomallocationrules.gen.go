package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReorderCustomAllocationRulesCmd = &cobra.Command{
	Use:   "reordercustomallocationrules",
	Short: "Reorder custom allocation rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cost/arbitrary_rule/reorder")
		fmt.Println("OperationID: ReorderCustomAllocationRules")
	},
}

func init() {
	Cmd.AddCommand(ReorderCustomAllocationRulesCmd)
}
