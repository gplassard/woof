package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchIssuesCmd = &cobra.Command{
	Use: "search-issues",

	Short: "Search error tracking issues",
	Long: `Search error tracking issues
Documentation: https://docs.datadoghq.com/api/latest/error-tracking/#search-issues`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IssuesSearchResponse
		var err error

		optionalParams := datadogV2.NewSearchIssuesOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchIssues(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-issues")

		cmd.Println(cmdutil.FormatJSON(res, "error_tracking_search_result"))
	},
}

func init() {

	SearchIssuesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchIssuesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	SearchIssuesCmd.Flags().String("include", "", "Comma-separated list of relationship objects that should be included in the response. Possible values are 'issue', 'issue.assignee', 'issue.case', and 'issue.team_owners'.")

	Cmd.AddCommand(SearchIssuesCmd)
}
