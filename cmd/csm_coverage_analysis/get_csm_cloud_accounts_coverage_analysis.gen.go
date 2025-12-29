package csm_coverage_analysis

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCSMCloudAccountsCoverageAnalysisCmd = &cobra.Command{
	Use:   "get-csm-cloud-accounts-coverage-analysis",
	
	Short: "Get the CSM Cloud Accounts Coverage Analysis",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		res, _, err := api.GetCSMCloudAccountsCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-csm-cloud-accounts-coverage-analysis")

		cmd.Println(cmdutil.FormatJSON(res, "csm_coverage_analysis"))
	},
}

func init() {
	Cmd.AddCommand(GetCSMCloudAccountsCoverageAnalysisCmd)
}
