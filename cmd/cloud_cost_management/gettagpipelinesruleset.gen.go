package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "gettagpipelinesruleset",
	Short: "Get a tag pipeline ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/tags/enrichment/{ruleset_id}")
		fmt.Println("OperationID: GetTagPipelinesRuleset")
	},
}

func init() {
	Cmd.AddCommand(GetTagPipelinesRulesetCmd)
}
