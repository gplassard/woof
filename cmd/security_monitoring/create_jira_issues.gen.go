package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateJiraIssuesCmd = &cobra.Command{
	Use:   "create-jira-issues",
	
	Short: "Create Jira issues for security findings",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateJiraIssues(client.NewContext(apiKey, appKey, site), datadogV2.CreateJiraIssueRequestArray{})
		if err != nil {
			log.Fatalf("failed to create-jira-issues: %v", err)
		}

		cmdutil.PrintJSON(res, "cases")
	},
}

func init() {
	Cmd.AddCommand(CreateJiraIssuesCmd)
}
