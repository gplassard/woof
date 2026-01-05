package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateJiraIssuesCmd = &cobra.Command{
	Use: "create-jira-issues",

	Short: "Create Jira issues for security findings",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateJiraIssues(client.NewContext(apiKey, appKey, site), datadogV2.CreateJiraIssueRequestArray{})
		cmdutil.HandleError(err, "failed to create-jira-issues")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(CreateJiraIssuesCmd)
}
