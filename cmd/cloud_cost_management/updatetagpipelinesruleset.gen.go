package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "updatetagpipelinesruleset",
	Short: "Update tag pipeline ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/tags/enrichment/{ruleset_id}")
		fmt.Println("OperationID: UpdateTagPipelinesRuleset")
	},
}

func init() {
	Cmd.AddCommand(UpdateTagPipelinesRulesetCmd)
}
