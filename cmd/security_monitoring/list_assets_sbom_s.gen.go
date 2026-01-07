package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAssetsSBOMsCmd = &cobra.Command{
	Use: "list-assets-sbom-s",

	Short: "List assets SBOMs",
	Long: `List assets SBOMs
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-assets-sbom-s`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListAssetsSBOMsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAssetsSBOMs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-assets-sbom-s")

		fmt.Println(cmdutil.FormatJSON(res, "sboms"))
	},
}

func init() {

	Cmd.AddCommand(ListAssetsSBOMsCmd)
}
