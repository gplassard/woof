package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateScorecardOutcomesAsyncCmd = &cobra.Command{
	Use:   "updatescorecardoutcomesasync",
	Short: "Update Scorecard outcomes asynchronously",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/scorecard/outcomes")
		fmt.Println("OperationID: UpdateScorecardOutcomesAsync")
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardOutcomesAsyncCmd)
}
