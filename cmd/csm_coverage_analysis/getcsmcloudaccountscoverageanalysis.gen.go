package csm_coverage_analysis

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCSMCloudAccountsCoverageAnalysisCmd = &cobra.Command{
	Use:   "getcsmcloudaccountscoverageanalysis",
	Short: "Get the CSM Cloud Accounts Coverage Analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/csm/onboarding/coverage_analysis/cloud_accounts")
		fmt.Println("OperationID: GetCSMCloudAccountsCoverageAnalysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMCloudAccountsCoverageAnalysisCmd)
}
