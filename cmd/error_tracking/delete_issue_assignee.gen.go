package error_tracking

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIssueAssigneeCmd = &cobra.Command{
	Use:   "delete-issue-assignee [issue_id]",
	
	Short: "Remove the assignee of an issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		_, err := api.DeleteIssueAssignee(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-issue-assignee: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIssueAssigneeCmd)
}
