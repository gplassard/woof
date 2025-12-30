package error_tracking

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIssueCmd = &cobra.Command{
	Use: "get-issue [issue_id]",

	Short: "Get the details of an error tracking issue",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		res, _, err := api.GetIssue(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-issue")

		cmd.Println(cmdutil.FormatJSON(res, "issue"))
	},
}

func init() {
	Cmd.AddCommand(GetIssueCmd)
}
