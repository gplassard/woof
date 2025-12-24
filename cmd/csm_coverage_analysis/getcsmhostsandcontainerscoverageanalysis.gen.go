package csm_coverage_analysis

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCSMHostsAndContainersCoverageAnalysisCmd = &cobra.Command{
	Use:   "getcsmhostsandcontainerscoverageanalysis",
	Short: "Get the CSM Hosts and Containers Coverage Analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/csm/onboarding/coverage_analysis/hosts_and_containers")
		fmt.Println("OperationID: GetCSMHostsAndContainersCoverageAnalysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMHostsAndContainersCoverageAnalysisCmd)
}
