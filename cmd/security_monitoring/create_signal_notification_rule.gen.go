package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSignalNotificationRuleCmd = &cobra.Command{
	Use: "create-signal-notification-rule",

	Short: "Create a new signal-based notification rule",
	Long: `Create a new signal-based notification rule
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-signal-notification-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.NotificationRuleResponse
		var err error

		var body datadogV2.CreateNotificationRuleParameters
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSignalNotificationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-signal-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "signal_notification_rule"))
	},
}

func init() {

	CreateSignalNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSignalNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSignalNotificationRuleCmd)
}
