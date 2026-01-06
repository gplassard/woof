package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMonitorNotificationRulesCmd = &cobra.Command{
	Use:     "get-monitor-notification-rules",
	Aliases: []string{"get-notification-rules"},
	Short:   "Get all monitor notification rules",
	Long: `Get all monitor notification rules
Documentation: https://docs.datadoghq.com/api/latest/monitors/#get-monitor-notification-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleListResponse
		var err error

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonitorNotificationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-monitor-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {

	Cmd.AddCommand(GetMonitorNotificationRulesCmd)
}
