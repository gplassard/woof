package csm_coverage_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCSMServerlessCoverageAnalysisCmd = &cobra.Command{
	Use: "get-csm-serverless-coverage-analysis",

	Short: "Get the CSM Serverless Coverage Analysis",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CsmServerlessCoverageAnalysisResponse
		var err error

		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		res, _, err = api.GetCSMServerlessCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-csm-serverless-coverage-analysis")

		cmd.Println(cmdutil.FormatJSON(res, "csm_coverage_analysis"))
	},
}

func init() {
	Cmd.AddCommand(GetCSMServerlessCoverageAnalysisCmd)
}
