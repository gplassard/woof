package error_tracking

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIssueAssigneeCmd = &cobra.Command{
	Use:   "update-issue-assignee [issue_id]",
	
	Short: "Update the assignee of an issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.UpdateIssueAssignee(client.NewContext(apiKey, appKey, site), args[0], datadogV2.IssueUpdateAssigneeRequest{})
		if err != nil {
			log.Fatalf("failed to update-issue-assignee: %v", err)
		}

		cmdutil.PrintJSON(res, "issue")
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueAssigneeCmd)
}
