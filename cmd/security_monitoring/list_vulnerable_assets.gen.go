package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use: "list-vulnerable-assets",

	Short: "List vulnerable assets",
	Long: `List vulnerable assets
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-vulnerable-assets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListVulnerableAssetsResponse
		var err error

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListVulnerableAssets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-vulnerable-assets")

		fmt.Println(cmdutil.FormatJSON(res, "vulnerable_asset"))
	},
}

func init() {

	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
