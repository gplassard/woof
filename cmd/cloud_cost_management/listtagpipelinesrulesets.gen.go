package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTagPipelinesRulesetsCmd = &cobra.Command{
	Use:   "listtagpipelinesrulesets",
	Short: "List tag pipeline rulesets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/tags/enrichment")
		fmt.Println("OperationID: ListTagPipelinesRulesets")
	},
}

func init() {
	Cmd.AddCommand(ListTagPipelinesRulesetsCmd)
}
