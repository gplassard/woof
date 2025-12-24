package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AttachJiraIssueCmd = &cobra.Command{
	Use:   "attachjiraissue",
	Short: "Attach security findings to a Jira issue",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security/findings/jira_issues")
		fmt.Println("OperationID: AttachJiraIssue")
	},
}

func init() {
	Cmd.AddCommand(AttachJiraIssueCmd)
}
