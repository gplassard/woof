package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteProjectNotificationRuleCmd = &cobra.Command{
	Use: "delete-project-notification-rule [project_id] [notification_rule_id]",

	Short: "Delete a notification rule",
	Long: `Delete a notification rule
Documentation: https://docs.datadoghq.com/api/latest/case-management/#delete-project-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteProjectNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-project-notification-rule")

	},
}

func init() {

	Cmd.AddCommand(DeleteProjectNotificationRuleCmd)
}
