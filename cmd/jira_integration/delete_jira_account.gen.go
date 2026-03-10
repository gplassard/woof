package jira_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteJiraAccountCmd = &cobra.Command{
	Use: "delete-jira-account [account_id]",

	Short: "Delete Jira account",
	Long: `Delete Jira account
Documentation: https://docs.datadoghq.com/api/latest/jira-integration/#delete-jira-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewJiraIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteJiraAccount(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-jira-account")

	},
}

func init() {

	Cmd.AddCommand(DeleteJiraAccountCmd)
}
