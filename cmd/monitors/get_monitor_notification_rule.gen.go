package monitors

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "get-monitor-notification-rule [rule_id]",
	Aliases: []string{ "get-notification-rule", },
	Short: "Get a monitor notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-monitor-notification-rule")

		cmdutil.PrintJSON(res, "monitor-notification-rule")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorNotificationRuleCmd)
}
