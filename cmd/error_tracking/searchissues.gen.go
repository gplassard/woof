package error_tracking

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchIssuesCmd = &cobra.Command{
	Use:   "searchissues",
	Short: "Search error tracking issues",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.SearchIssues(client.NewContext(apiKey, appKey, site), datadogV2.IssuesSearchRequest{})
		if err != nil {
			log.Fatalf("failed to searchissues: %v", err)
		}

		cmdutil.PrintJSON(res, "error_tracking")
	},
}

func init() {
	Cmd.AddCommand(SearchIssuesCmd)
}
