package jira_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListJiraIssueTemplatesCmd = &cobra.Command{
	Use: "list-jira-issue-templates",

	Short: "List Jira issue templates",
	Long: `List Jira issue templates
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#list-jira-issue-templates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JiraIssueTemplatesResponse
		var err error

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListJiraIssueTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-jira-issue-templates")

		fmt.Println(cmdutil.FormatJSON(res, "jira-issue-template"))
	},
}

func init() {

	Cmd.AddCommand(ListJiraIssueTemplatesCmd)
}
