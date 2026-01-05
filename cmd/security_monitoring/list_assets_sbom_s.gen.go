package security_monitoring

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAssetsSBOMsCmd = &cobra.Command{
	Use: "list-assets-sbom-s",

	Short: "List assets SBOMs",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListAssetsSBOMs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-assets-sbom-s")

		cmd.Println(cmdutil.FormatJSON(res, "sboms"))
	},
}

func init() {
	Cmd.AddCommand(ListAssetsSBOMsCmd)
}
