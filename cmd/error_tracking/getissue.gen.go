package error_tracking

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetIssueCmd = &cobra.Command{
	Use:   "getissue [issue_id]",
	Short: "Get the details of an error tracking issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.GetIssue(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getissue: %v", err)
		}

		cmdutil.PrintJSON(res, "error_tracking")
	},
}

func init() {
	Cmd.AddCommand(GetIssueCmd)
}
