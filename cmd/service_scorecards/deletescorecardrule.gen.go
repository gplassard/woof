package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteScorecardRuleCmd = &cobra.Command{
	Use:   "deletescorecardrule",
	Short: "Delete a rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/scorecard/rules/{rule_id}")
		fmt.Println("OperationID: DeleteScorecardRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteScorecardRuleCmd)
}
