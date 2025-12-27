package csm_coverage_analysis

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMCloudAccountsCoverageAnalysisCmd = &cobra.Command{
	Use:   "get_c_s_m_cloud_accounts_coverage_analysis",
	Short: "Get the CSM Cloud Accounts Coverage Analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		res, _, err := api.GetCSMCloudAccountsCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_c_s_m_cloud_accounts_coverage_analysis: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_coverage_analysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMCloudAccountsCoverageAnalysisCmd)
}
