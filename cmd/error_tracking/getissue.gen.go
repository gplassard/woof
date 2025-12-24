package error_tracking

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIssueCmd = &cobra.Command{
	Use:   "getissue",
	Short: "Get the details of an error tracking issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/error-tracking/issues/{issue_id}")
		fmt.Println("OperationID: GetIssue")
	},
}

func init() {
	Cmd.AddCommand(GetIssueCmd)
}
