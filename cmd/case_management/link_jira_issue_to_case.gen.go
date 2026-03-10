package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var LinkJiraIssueToCaseCmd = &cobra.Command{
	Use: "link-jira-issue-to-case [case_id]",

	Short: "Link existing Jira issue to case",
	Long: `Link existing Jira issue to case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#link-jira-issue-to-case`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.JiraIssueLinkRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.LinkJiraIssueToCase(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to link-jira-issue-to-case")

	},
}

func init() {

	LinkJiraIssueToCaseCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	LinkJiraIssueToCaseCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(LinkJiraIssueToCaseCmd)
}
