package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSignalNotificationRulesCmd = &cobra.Command{
	Use:   "get-signal-notification-rules",
	Short: "Get the list of signal-based notification rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSignalNotificationRules(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-signal-notification-rules: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(GetSignalNotificationRulesCmd)
}
