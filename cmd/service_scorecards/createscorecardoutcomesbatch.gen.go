package service_scorecards

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateScorecardOutcomesBatchCmd = &cobra.Command{
	Use:   "createscorecardoutcomesbatch",
	Short: "Create outcomes batch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/scorecard/outcomes/batch")
		fmt.Println("OperationID: CreateScorecardOutcomesBatch")
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardOutcomesBatchCmd)
}
