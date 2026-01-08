package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteIssueAssigneeCmd = &cobra.Command{
	Use: "delete-issue-assignee [issue_id]",

	Short: "Remove the assignee of an issue",
	Long: `Remove the assignee of an issue
Documentation: https://docs.datadoghq.com/api/latest/error-tracking/#delete-issue-assignee`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteIssueAssignee(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-issue-assignee")

	},
}

func init() {

	Cmd.AddCommand(DeleteIssueAssigneeCmd)
}
