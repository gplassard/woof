package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateJiraIssuesCmd = &cobra.Command{
	Use: "create-jira-issues [payload]",

	Short: "Create Jira issues for security findings",
	Long: `Create Jira issues for security findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-jira-issues`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FindingCaseResponseArray
		var err error

		var body datadogV2.CreateJiraIssueRequestArray
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.CreateJiraIssues(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-jira-issues")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(CreateJiraIssuesCmd)
}
