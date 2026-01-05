package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIssueStateCmd = &cobra.Command{
	Use: "update-issue-state [issue_id] [payload]",

	Short: "Update the state of an issue",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IssueResponse
		var err error

		var body datadogV2.IssueUpdateStateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err = api.UpdateIssueState(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-issue-state")

		cmd.Println(cmdutil.FormatJSON(res, "issue"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIssueStateCmd)
}
