package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityMonitoringCriticalAssetsCmd = &cobra.Command{
	Use:     "list-security-monitoring-critical-assets",
	Aliases: []string{"list-critical-assets"},
	Short:   "Get all critical assets",
	Long: `Get all critical assets
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-security-monitoring-critical-assets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringCriticalAssetsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSecurityMonitoringCriticalAssets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-monitoring-critical-assets")

		fmt.Println(cmdutil.FormatJSON(res, "critical_assets"))
	},
}

func init() {

	Cmd.AddCommand(ListSecurityMonitoringCriticalAssetsCmd)
}
