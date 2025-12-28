package csm_coverage_analysis

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMHostsAndContainersCoverageAnalysisCmd = &cobra.Command{
	Use:   "get-c-s-m-hosts-and-containers-coverage-analysis",
	Short: "Get the CSM Hosts and Containers Coverage Analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		res, _, err := api.GetCSMHostsAndContainersCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-c-s-m-hosts-and-containers-coverage-analysis: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_coverage_analysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMHostsAndContainersCoverageAnalysisCmd)
}
