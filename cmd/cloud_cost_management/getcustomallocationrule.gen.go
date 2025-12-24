package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCustomAllocationRuleCmd = &cobra.Command{
	Use:   "getcustomallocationrule",
	Short: "Get custom allocation rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/arbitrary_rule/{rule_id}")
		fmt.Println("OperationID: GetCustomAllocationRule")
	},
}

func init() {
	Cmd.AddCommand(GetCustomAllocationRuleCmd)
}
