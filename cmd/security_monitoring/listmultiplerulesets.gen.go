package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMultipleRulesetsCmd = &cobra.Command{
	Use:   "listmultiplerulesets",
	Short: "Ruleset get multiple",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/static-analysis/rulesets")
		fmt.Println("OperationID: ListMultipleRulesets")
	},
}

func init() {
	Cmd.AddCommand(ListMultipleRulesetsCmd)
}
