package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIssueAssigneeCmd = &cobra.Command{
	Use: "update-issue-assignee [issue_id]",

	Short: "Update the assignee of an issue",
	Long: `Update the assignee of an issue
Documentation: https://docs.datadoghq.com/api/latest/error-tracking/#update-issue-assignee`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IssueResponse
		var err error

		var body datadogV2.IssueUpdateAssigneeRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err = api.UpdateIssueAssignee(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-issue-assignee")

		cmd.Println(cmdutil.FormatJSON(res, "issue"))
	},
}

func init() {

	UpdateIssueAssigneeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIssueAssigneeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIssueAssigneeCmd)
}
