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

var UpdateJiraIssueTemplateCmd = &cobra.Command{
	Use: "update-jira-issue-template [issue_template_id]",

	Short: "Update Jira issue template",
	Long: `Update Jira issue template
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#update-jira-issue-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JiraIssueTemplateResponse
		var err error

		var body datadogV2.JiraIssueTemplateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateJiraIssueTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-jira-issue-template")

		fmt.Println(cmdutil.FormatJSON(res, "jira-issue-template"))
	},
}

func init() {

	UpdateJiraIssueTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateJiraIssueTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateJiraIssueTemplateCmd)
}
