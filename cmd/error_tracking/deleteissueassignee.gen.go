package error_tracking

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIssueAssigneeCmd = &cobra.Command{
	Use:   "deleteissueassignee",
	Short: "Remove the assignee of an issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/error-tracking/issues/{issue_id}/assignee")
		fmt.Println("OperationID: DeleteIssueAssignee")
	},
}

func init() {
	Cmd.AddCommand(DeleteIssueAssigneeCmd)
}
