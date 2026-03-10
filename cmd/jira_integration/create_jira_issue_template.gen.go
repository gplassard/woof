package jira_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateJiraIssueTemplateCmd = &cobra.Command{
	Use: "create-jira-issue-template",

	Short: "Create Jira issue template",
	Long: `Create Jira issue template
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#create-jira-issue-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JiraIssueTemplateResponse
		var err error

		var body datadogV2.JiraIssueTemplateCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateJiraIssueTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-jira-issue-template")

		fmt.Println(cmdutil.FormatJSON(res, "jira-issue-template"))
	},
}

func init() {

	CreateJiraIssueTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateJiraIssueTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateJiraIssueTemplateCmd)
}
