package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSignalNotificationRulesCmd = &cobra.Command{
	Use: "get-signal-notification-rules",

	Short: "Get the list of signal-based notification rules",
	Long: `Get the list of signal-based notification rules
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-signal-notification-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetSignalNotificationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-signal-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	Cmd.AddCommand(GetSignalNotificationRulesCmd)
}
