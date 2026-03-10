package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSecurityMonitoringCriticalAssetCmd = &cobra.Command{
	Use:     "create-security-monitoring-critical-asset",
	Aliases: []string{"create-critical-asset"},
	Short:   "Create a critical asset",
	Long: `Create a critical asset
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#create-security-monitoring-critical-asset`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityMonitoringCriticalAssetResponse
		var err error

		var body datadogV2.SecurityMonitoringCriticalAssetCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSecurityMonitoringCriticalAsset(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-monitoring-critical-asset")

		fmt.Println(cmdutil.FormatJSON(res, "critical_assets"))
	},
}

func init() {

	CreateSecurityMonitoringCriticalAssetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSecurityMonitoringCriticalAssetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSecurityMonitoringCriticalAssetCmd)
}
