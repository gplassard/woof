package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UnlinkJiraIssueCmd = &cobra.Command{
	Use: "unlink-jira-issue [case_id]",

	Short: "Remove Jira issue link from case",
	Long: `Remove Jira issue link from case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#unlink-jira-issue`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UnlinkJiraIssue(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to unlink-jira-issue")

	},
}

func init() {

	Cmd.AddCommand(UnlinkJiraIssueCmd)
}
