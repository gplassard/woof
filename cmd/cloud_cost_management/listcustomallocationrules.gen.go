package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCustomAllocationRulesCmd = &cobra.Command{
	Use:   "listcustomallocationrules",
	Short: "List custom allocation rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/arbitrary_rule")
		fmt.Println("OperationID: ListCustomAllocationRules")
	},
}

func init() {
	Cmd.AddCommand(ListCustomAllocationRulesCmd)
}
