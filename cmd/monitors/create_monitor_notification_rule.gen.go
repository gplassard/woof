package monitors

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "create-monitor-notification-rule",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create a monitor notification rule",
	Long: `Create a monitor notification rule
Documentation: https://docs.datadoghq.com/api/latest/monitors/#create-monitor-notification-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleResponse
		var err error

		var body datadogV2.MonitorNotificationRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateMonitorNotificationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-monitor-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {

	CreateMonitorNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateMonitorNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateMonitorNotificationRuleCmd)
}
