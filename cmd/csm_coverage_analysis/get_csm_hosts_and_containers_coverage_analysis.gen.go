package csm_coverage_analysis

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCSMHostsAndContainersCoverageAnalysisCmd = &cobra.Command{
	Use: "get-csm-hosts-and-containers-coverage-analysis",

	Short: "Get the CSM Hosts and Containers Coverage Analysis",
	Long: `Get the CSM Hosts and Containers Coverage Analysis
Documentation: https://docs.datadoghq.com/api/latest/c-s-m-coverage-analysis/#get-csm-hosts-and-containers-coverage-analysis`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CsmHostsAndContainersCoverageAnalysisResponse
		var err error

		api := datadogV2.NewCSMCoverageAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCSMHostsAndContainersCoverageAnalysis(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-csm-hosts-and-containers-coverage-analysis")

		fmt.Println(cmdutil.FormatJSON(res, "c_s_m_hosts_and_containers_coverage_analysi"))
	},
}

func init() {

	Cmd.AddCommand(GetCSMHostsAndContainersCoverageAnalysisCmd)
}
