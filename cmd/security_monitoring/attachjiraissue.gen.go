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
	Use:   "attachjiraissue",
	Short: "Attach security findings to a Jira issue",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.AttachJiraIssue(client.NewContext(apiKey, appKey, site), datadogV2.AttachJiraIssueRequest{})
		if err != nil {
			log.Fatalf("failed to attachjiraissue: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(AttachJiraIssueCmd)
}
