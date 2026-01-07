package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AttachJiraIssueCmd = &cobra.Command{
	Use: "attach-jira-issue",

	Short: "Attach security findings to a Jira issue",
	Long: `Attach security findings to a Jira issue
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#attach-jira-issue`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FindingCaseResponse
		var err error

		var body datadogV2.AttachJiraIssueRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AttachJiraIssue(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to attach-jira-issue")

		fmt.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {

	AttachJiraIssueCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AttachJiraIssueCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AttachJiraIssueCmd)
}
