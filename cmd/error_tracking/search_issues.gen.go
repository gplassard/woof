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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.SearchIssues(client.NewContext(apiKey, appKey, site), datadogV2.IssuesSearchRequest{})
		cmdutil.HandleError(err, "failed to search-issues")

		cmd.Println(cmdutil.FormatJSON(res, "error_tracking_search_result"))
	},
}

func init() {
	Cmd.AddCommand(SearchIssuesCmd)
}
