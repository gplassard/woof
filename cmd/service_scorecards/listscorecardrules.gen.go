package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListScorecardRulesCmd = &cobra.Command{
	Use:   "listscorecardrules",
	Short: "List all rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/scorecard/rules")
		fmt.Println("OperationID: ListScorecardRules")
	},
}

func init() {
	Cmd.AddCommand(ListScorecardRulesCmd)
}
