package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AttachJiraIssueCmd = &cobra.Command{
	Use: "attach-jira-issue",

	Short: "Attach security findings to a Jira issue",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.AttachJiraIssue(client.NewContext(apiKey, appKey, site), datadogV2.AttachJiraIssueRequest{})
		cmdutil.HandleError(err, "failed to attach-jira-issue")

		cmd.Println(cmdutil.FormatJSON(res, "cases"))
	},
}

func init() {
	Cmd.AddCommand(AttachJiraIssueCmd)
}
