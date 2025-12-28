package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "create-monitor-notification-rule",
	Aliases: []string{ "create-notification-rule", },
	Short: "Create a monitor notification rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.CreateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), datadogV2.MonitorNotificationRuleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-monitor-notification-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "monitor-notification-rule")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorNotificationRuleCmd)
}
