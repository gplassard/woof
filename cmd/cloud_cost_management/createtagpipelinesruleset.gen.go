package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "createtagpipelinesruleset",
	Short: "Create tag pipeline ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/tags/enrichment")
		fmt.Println("OperationID: CreateTagPipelinesRuleset")
	},
}

func init() {
	Cmd.AddCommand(CreateTagPipelinesRulesetCmd)
}
