package error_tracking

import (
	"fmt"
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

		var body datadogV2.IssuesSearchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchIssues(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-issues")

		fmt.Println(cmdutil.FormatJSON(res, "error_tracking_search_result"))
	},
}

func init() {

	SearchIssuesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchIssuesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchIssuesCmd)
}
