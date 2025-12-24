package error_tracking

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIssueStateCmd = &cobra.Command{
	Use:   "updateissuestate [issue_id]",
	Short: "Update the state of an issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.UpdateIssueState(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IssueUpdateStateRequest{})
		if err != nil {
			log.Fatalf("failed to updateissuestate: %v", err)
		}

		cmdutil.PrintJSON(res, "error_tracking")
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueStateCmd)
}
