package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCaseJiraIssueCmd = &cobra.Command{
	Use: "create-case-jira-issue [case_id]",

	Short: "Create Jira issue for case",
	Long: `Create Jira issue for case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-case-jira-issue`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.JiraIssueCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.CreateCaseJiraIssue(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-case-jira-issue")

	},
}

func init() {

	CreateCaseJiraIssueCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCaseJiraIssueCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCaseJiraIssueCmd)
}
