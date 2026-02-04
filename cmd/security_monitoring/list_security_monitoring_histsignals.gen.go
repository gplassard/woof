package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:     "list-security-monitoring-histsignals",
	Aliases: []string{"list-histsignals"},
	Short:   "List hist signals",
	Long: `List hist signals
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-histsignals`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalsListResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-histsignals")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring_histsignal"))
	},
}

func init() {

	Cmd.AddCommand(ListSecurityMonitoringHistsignalsCmd)
}
