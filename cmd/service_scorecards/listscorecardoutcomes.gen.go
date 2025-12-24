package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListScorecardOutcomesCmd = &cobra.Command{
	Use:   "listscorecardoutcomes",
	Short: "List all rule outcomes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/scorecard/outcomes")
		fmt.Println("OperationID: ListScorecardOutcomes")
	},
}

func init() {
	Cmd.AddCommand(ListScorecardOutcomesCmd)
}
