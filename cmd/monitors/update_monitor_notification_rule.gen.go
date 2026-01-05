package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "update-monitor-notification-rule [rule_id]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update a monitor notification rule",
	Long: `Update a monitor notification rule
Documentation: https://docs.datadoghq.com/api/latest/monitors/#update-monitor-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleResponse
		var err error

		var body datadogV2.MonitorNotificationRuleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.UpdateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-monitor-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {

	UpdateMonitorNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateMonitorNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateMonitorNotificationRuleCmd)
}
