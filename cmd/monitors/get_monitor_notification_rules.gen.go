package monitors

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMonitorNotificationRulesCmd = &cobra.Command{
	Use:   "get-monitor-notification-rules",
	Aliases: []string{ "get-notification-rules", },
	Short: "Get all monitor notification rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorNotificationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-monitor-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {
	Cmd.AddCommand(GetMonitorNotificationRulesCmd)
}
