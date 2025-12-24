package csm_coverage_analysis

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMCloudAccountsCoverageAnalysisCmd = &cobra.Command{
	Use:   "getcsmcloudaccountscoverageanalysis",
	Short: "Get the CSM Cloud Accounts Coverage Analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		res, _, err := api.GetCSMCloudAccountsCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getcsmcloudaccountscoverageanalysis: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_coverage_analysis")
	},
}

func init() {
	Cmd.AddCommand(GetCSMCloudAccountsCoverageAnalysisCmd)
}
