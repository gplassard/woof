package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringSignalCmd = &cobra.Command{
	Use:     "get-security-monitoring-signal [signal_id]",
	Aliases: []string{"get-signal"},
	Short:   "Get a signal's details",
	Long: `Get a signal's details
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-signal`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecurityMonitoringSignal(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-signal")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring_signal"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityMonitoringSignalCmd)
}
