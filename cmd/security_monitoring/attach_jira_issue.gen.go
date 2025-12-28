package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AttachJiraIssueCmd = &cobra.Command{
	Use:   "attach_jira_issue",
	Short: "Attach security findings to a Jira issue",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.AttachJiraIssue(client.NewContext(apiKey, appKey, site), datadogV2.AttachJiraIssueRequest{})
		if err != nil {
			log.Fatalf("failed to attach_jira_issue: %v", err)
		}

		cmdutil.PrintJSON(res, "cases")
	},
}

func init() {
	Cmd.AddCommand(AttachJiraIssueCmd)
}
