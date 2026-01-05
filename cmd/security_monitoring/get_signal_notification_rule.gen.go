package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSignalNotificationRuleCmd = &cobra.Command{
	Use: "get-signal-notification-rule [id]",

	Short: "Get details of a signal-based notification rule",
	Long: `Get details of a signal-based notification rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-signal-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.NotificationRuleResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetSignalNotificationRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-signal-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(GetSignalNotificationRuleCmd)
}
