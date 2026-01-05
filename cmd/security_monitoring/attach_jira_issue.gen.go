package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AttachJiraIssueCmd = &cobra.Command{
	Use: "attach-jira-issue [payload]",

	Short: "Attach security findings to a Jira issue",
	Long: `Attach security findings to a Jira issue
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#attach-jira-issue`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FindingCaseResponse
		var err error

		var body datadogV2.AttachJiraIssueRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.AttachJiraIssue(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to attach-jira-issue")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(AttachJiraIssueCmd)
}
