package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchIssuesCmd = &cobra.Command{
	Use: "search-issues [payload]",

	Short: "Search error tracking issues",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.IssuesSearchRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.SearchIssues(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-issues")

		cmd.Println(cmdutil.FormatJSON(res, "error_tracking_search_result"))
	},
}

func init() {
	Cmd.AddCommand(SearchIssuesCmd)
}
