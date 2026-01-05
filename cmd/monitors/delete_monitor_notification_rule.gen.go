package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "delete-monitor-notification-rule [rule_id]",
	Aliases: []string{"delete-notification-rule"},
	Short:   "Delete a monitor notification rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.DeleteMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-monitor-notification-rule")

	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorNotificationRuleCmd)
}
