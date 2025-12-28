package error_tracking

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIssueStateCmd = &cobra.Command{
	Use:   "update_issue_state [issue_id]",
	Short: "Update the state of an issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.UpdateIssueState(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IssueUpdateStateRequest{})
		if err != nil {
			log.Fatalf("failed to update_issue_state: %v", err)
		}

		cmdutil.PrintJSON(res, "issue")
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueStateCmd)
}
