package monitors

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetMonitorNotificationRuleCmd = &cobra.Command{
	Use:     "get-monitor-notification-rule [rule_id]",
	Aliases: []string{"get-notification-rule"},
	Short:   "Get a monitor notification rule",
	Long: `Get a monitor notification rule
Documentation: https://docs.datadoghq.com/api/latest/monitors/#get-monitor-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MonitorNotificationRuleResponse
		var err error

		optionalParams := datadogV2.NewGetMonitorNotificationRuleOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMonitorNotificationRule(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-monitor-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "monitor-notification-rule"))
	},
}

func init() {

	GetMonitorNotificationRuleCmd.Flags().String("include", "", "Comma-separated list of resource paths for related resources to include in the response. Supported resource path is 'created_by'.")

	Cmd.AddCommand(GetMonitorNotificationRuleCmd)
}
