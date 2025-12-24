package error_tracking

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIssueStateCmd = &cobra.Command{
	Use:   "updateissuestate",
	Short: "Update the state of an issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/error-tracking/issues/{issue_id}/state")
		fmt.Println("OperationID: UpdateIssueState")
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueStateCmd)
}
