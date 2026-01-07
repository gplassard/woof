package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScannedAssetsMetadataCmd = &cobra.Command{
	Use: "list-scanned-assets-metadata",

	Short: "List scanned assets metadata",
	Long: `List scanned assets metadata
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-scanned-assets-metadata`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ScannedAssetsMetadata
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScannedAssetsMetadata(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scanned-assets-metadata")

		fmt.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	Cmd.AddCommand(ListScannedAssetsMetadataCmd)
}
