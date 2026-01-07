package error_tracking

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetIssueCmd = &cobra.Command{
	Use: "get-issue [issue_id]",

	Short: "Get the details of an error tracking issue",
	Long: `Get the details of an error tracking issue
Documentation: https://docs.datadoghq.com/api/latest/error-tracking/#get-issue`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IssueResponse
		var err error

		optionalParams := datadogV2.NewGetIssueOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewErrorTrackingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetIssue(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-issue")

		cmd.Println(cmdutil.FormatJSON(res, "issue"))
	},
}

func init() {

	GetIssueCmd.Flags().String("include", "", "Comma-separated list of relationship objects that should be included in the response. Possible values are 'assignee', 'case', and 'team_owners'.")

	Cmd.AddCommand(GetIssueCmd)
}
