package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCustomAllocationRuleCmd = &cobra.Command{
	Use:   "updatecustomallocationrule",
	Short: "Update custom allocation rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/cost/arbitrary_rule/{rule_id}")
		fmt.Println("OperationID: UpdateCustomAllocationRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomAllocationRuleCmd)
}
