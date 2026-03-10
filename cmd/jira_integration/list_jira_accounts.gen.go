package jira_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListJiraAccountsCmd = &cobra.Command{
	Use: "list-jira-accounts",

	Short: "List Jira accounts",
	Long: `List Jira accounts
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#list-jira-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.JiraAccountsResponse
		var err error

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListJiraAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-jira-accounts")

		fmt.Println(cmdutil.FormatJSON(res, "jira_integration"))
	},
}

func init() {

	Cmd.AddCommand(ListJiraAccountsCmd)
}
