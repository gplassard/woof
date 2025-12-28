package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var PatchSignalNotificationRuleCmd = &cobra.Command{
	Use:   "patch_signal_notification_rule [id]",
	Short: "Patch a signal-based notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.PatchSignalNotificationRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.PatchNotificationRuleParameters{})
		if err != nil {
			log.Fatalf("failed to patch_signal_notification_rule: %v", err)
		}

		cmdutil.PrintJSON(res, "notification_rules")
	},
}

func init() {
	Cmd.AddCommand(PatchSignalNotificationRuleCmd)
}
