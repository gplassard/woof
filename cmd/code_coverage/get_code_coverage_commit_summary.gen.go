package code_coverage

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCodeCoverageCommitSummaryCmd = &cobra.Command{
	Use:     "get-code-coverage-commit-summary",
	Aliases: []string{"get-commit-summary"},
	Short:   "Get code coverage summary for a commit",
	Long: `Get code coverage summary for a commit
Documentation: https://docs.datadoghq.com/api/latest/code-coverage/#get-code-coverage-commit-summary`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CoverageSummaryResponse
		var err error

		var body datadogV2.CommitCoverageSummaryRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCodeCoverageApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCodeCoverageCommitSummary(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to get-code-coverage-commit-summary")

		fmt.Println(cmdutil.FormatJSON(res, "ci_app_coverage_summary"))
	},
}

func init() {

	GetCodeCoverageCommitSummaryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	GetCodeCoverageCommitSummaryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(GetCodeCoverageCommitSummaryCmd)
}
