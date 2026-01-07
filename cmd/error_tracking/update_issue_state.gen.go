package error_tracking

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateIssueStateCmd = &cobra.Command{
	Use: "update-issue-state [issue_id]",

	Short: "Update the state of an issue",
	Long: `Update the state of an issue
Documentation: https://docs.datadoghq.com/api/latest/error-tracking/#update-issue-state`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IssueResponse
		var err error

		var body datadogV2.IssueUpdateStateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIssueState(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-issue-state")

		fmt.Println(cmdutil.FormatJSON(res, "issue"))
	},
}

func init() {

	UpdateIssueStateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIssueStateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIssueStateCmd)
}
