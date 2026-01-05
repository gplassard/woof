package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSBOMCmd = &cobra.Command{
	Use: "get-sbom [asset_type] [filter[asset_name]]",

	Short: "Get SBOM",
	Long: `Get SBOM
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-sbom`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetSBOMResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetSBOM(client.NewContext(apiKey, appKey, site), datadogV2.AssetType(args[0]), args[1])
		cmdutil.HandleError(err, "failed to get-sbom")

		cmd.Println(cmdutil.FormatJSON(res, "sboms"))
	},
}

func init() {
	Cmd.AddCommand(GetSBOMCmd)
}
