package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var PatchSignalNotificationRuleCmd = &cobra.Command{
	Use: "patch-signal-notification-rule [id] [payload]",

	Short: "Patch a signal-based notification rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.NotificationRuleResponse
		var err error

		var body datadogV2.PatchNotificationRuleParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.PatchSignalNotificationRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to patch-signal-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(PatchSignalNotificationRuleCmd)
}
