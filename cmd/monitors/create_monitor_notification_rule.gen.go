package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "create-monitor-notification-rule [payload]",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create a monitor notification rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleResponse
		var err error

		var body datadogV2.MonitorNotificationRuleCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err = api.CreateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-monitor-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorNotificationRuleCmd)
}
