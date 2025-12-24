package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "deletetagpipelinesruleset",
	Short: "Delete tag pipeline ruleset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/tags/enrichment/{ruleset_id}")
		fmt.Println("OperationID: DeleteTagPipelinesRuleset")
	},
}

func init() {
	Cmd.AddCommand(DeleteTagPipelinesRulesetCmd)
}
