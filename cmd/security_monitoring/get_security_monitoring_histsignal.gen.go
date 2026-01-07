package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSecurityMonitoringHistsignalCmd = &cobra.Command{
	Use:     "get-security-monitoring-histsignal [histsignal_id]",
	Aliases: []string{"get-histsignal"},
	Short:   "Get a hist signal's details",
	Long: `Get a hist signal's details
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-security-monitoring-histsignal`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringSignalResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSecurityMonitoringHistsignal(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-security-monitoring-histsignal")

		fmt.Println(cmdutil.FormatJSON(res, "signal"))
	},
}

func init() {

	Cmd.AddCommand(GetSecurityMonitoringHistsignalCmd)
}
