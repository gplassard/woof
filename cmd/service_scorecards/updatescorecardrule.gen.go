package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateScorecardRuleCmd = &cobra.Command{
	Use:   "updatescorecardrule",
	Short: "Update an existing rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/scorecard/rules/{rule_id}")
		fmt.Println("OperationID: UpdateScorecardRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardRuleCmd)
}
