package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateScorecardRuleCmd = &cobra.Command{
	Use:   "createscorecardrule",
	Short: "Create a new rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/scorecard/rules")
		fmt.Println("OperationID: CreateScorecardRule")
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardRuleCmd)
}
