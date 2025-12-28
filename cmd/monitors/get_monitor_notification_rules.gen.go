package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMonitorNotificationRulesCmd = &cobra.Command{
	Use:   "get-monitor-notification-rules",
	Short: "Get all monitor notification rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorNotificationRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-monitor-notification-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-notification-rule")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorNotificationRulesCmd)
}
