package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringSignalsCmd = &cobra.Command{
	Use:     "list-security-monitoring-signals",
	Aliases: []string{"list-signals"},
	Short:   "Get a quick list of security signals",
	Long: `Get a quick list of security signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-signals`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringSignals(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-signals")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring_signal"))
	},
}

func init() {

	Cmd.AddCommand(ListSecurityMonitoringSignalsCmd)
}
