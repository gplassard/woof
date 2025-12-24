package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCustomAllocationRuleCmd = &cobra.Command{
	Use:   "createcustomallocationrule",
	Short: "Create custom allocation rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cost/arbitrary_rule")
		fmt.Println("OperationID: CreateCustomAllocationRule")
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAllocationRuleCmd)
}
