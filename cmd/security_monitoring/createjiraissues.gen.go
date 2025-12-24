package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateJiraIssuesCmd = &cobra.Command{
	Use:   "createjiraissues",
	Short: "Create Jira issues for security findings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security/findings/jira_issues")
		fmt.Println("OperationID: CreateJiraIssues")
	},
}

func init() {
	Cmd.AddCommand(CreateJiraIssuesCmd)
}
