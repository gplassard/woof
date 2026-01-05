package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "update-monitor-notification-rule [rule_id] [payload]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update a monitor notification rule",
	Long: `Update a monitor notification rule
Documentation: https://docs.datadoghq.com/api/latest/monitors/#update-monitor-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleResponse
		var err error

		var body datadogV2.MonitorNotificationRuleUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.UpdateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-monitor-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorNotificationRuleCmd)
}
