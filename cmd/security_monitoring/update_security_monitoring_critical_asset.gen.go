package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSecurityMonitoringCriticalAssetCmd = &cobra.Command{
	Use:     "update-security-monitoring-critical-asset [critical_asset_id]",
	Aliases: []string{"update-critical-asset"},
	Short:   "Update a critical asset",
	Long: `Update a critical asset
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-security-monitoring-critical-asset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringCriticalAssetResponse
		var err error

		var body datadogV2.SecurityMonitoringCriticalAssetUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateSecurityMonitoringCriticalAsset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-security-monitoring-critical-asset")

		fmt.Println(cmdutil.FormatJSON(res, "critical_assets"))
	},
}

func init() {

	UpdateSecurityMonitoringCriticalAssetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateSecurityMonitoringCriticalAssetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateSecurityMonitoringCriticalAssetCmd)
}
