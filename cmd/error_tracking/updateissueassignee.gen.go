package error_tracking

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIssueAssigneeCmd = &cobra.Command{
	Use:   "updateissueassignee",
	Short: "Update the assignee of an issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/error-tracking/issues/{issue_id}/assignee")
		fmt.Println("OperationID: UpdateIssueAssignee")
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueAssigneeCmd)
}
