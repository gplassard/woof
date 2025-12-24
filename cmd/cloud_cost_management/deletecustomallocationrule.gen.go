package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCustomAllocationRuleCmd = &cobra.Command{
	Use:   "deletecustomallocationrule",
	Short: "Delete custom allocation rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/arbitrary_rule/{rule_id}")
		fmt.Println("OperationID: DeleteCustomAllocationRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomAllocationRuleCmd)
}
