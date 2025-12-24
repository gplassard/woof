package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateJiraIssuesCmd = &cobra.Command{
	Use:   "createjiraissues",
	Short: "Create Jira issues for security findings",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateJiraIssues(client.NewContext(apiKey, appKey, site), datadogV2.CreateJiraIssueRequestArray{})
		if err != nil {
			log.Fatalf("failed to createjiraissues: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(CreateJiraIssuesCmd)
}
