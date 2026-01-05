package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSignalNotificationRuleCmd = &cobra.Command{
	Use: "create-signal-notification-rule",

	Short: "Create a new signal-based notification rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.CreateSignalNotificationRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateNotificationRuleParameters{})
		cmdutil.HandleError(err, "failed to create-signal-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(CreateSignalNotificationRuleCmd)
}
