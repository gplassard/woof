package csm_coverage_analysis

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCSMServerlessCoverageAnalysisCmd = &cobra.Command{
	Use:   "getcsmserverlesscoverageanalysis",
	Short: "Get the CSM Serverless Coverage Analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/csm/onboarding/coverage_analysis/serverless")
		fmt.Println("OperationID: GetCSMServerlessCoverageAnalysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMServerlessCoverageAnalysisCmd)
}
