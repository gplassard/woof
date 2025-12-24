package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReorderTagPipelinesRulesetsCmd = &cobra.Command{
	Use:   "reordertagpipelinesrulesets",
	Short: "Reorder tag pipeline rulesets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/tags/enrichment/reorder")
		fmt.Println("OperationID: ReorderTagPipelinesRulesets")
	},
}

func init() {
	Cmd.AddCommand(ReorderTagPipelinesRulesetsCmd)
}
