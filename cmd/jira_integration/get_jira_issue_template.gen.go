package jira_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetJiraIssueTemplateCmd = &cobra.Command{
	Use: "get-jira-issue-template [issue_template_id]",

	Short: "Get Jira issue template",
	Long: `Get Jira issue template
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#get-jira-issue-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JiraIssueTemplateResponse
		var err error

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetJiraIssueTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-jira-issue-template")

		fmt.Println(cmdutil.FormatJSON(res, "jira-issue-template"))
	},
}

func init() {

	Cmd.AddCommand(GetJiraIssueTemplateCmd)
}
