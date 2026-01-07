package security_monitoring

import (
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

		optionalParams := datadogV2.NewListScannedAssetsMetadataOptionalParameters()

		if cmd.Flags().Changed("page-token") {
			val, _ := cmd.Flags().GetString("page-token")
			optionalParams.WithPageToken(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("filter-asset.type") {
			val, _ := cmd.Flags().GetString("filter-asset.type")
			optionalParams.WithFilterAssetType(val)
		}

		if cmd.Flags().Changed("filter-asset.name") {
			val, _ := cmd.Flags().GetString("filter-asset.name")
			optionalParams.WithFilterAssetName(val)
		}

		if cmd.Flags().Changed("filter-last-success.origin") {
			val, _ := cmd.Flags().GetString("filter-last-success.origin")
			optionalParams.WithFilterLastSuccessOrigin(val)
		}

		if cmd.Flags().Changed("filter-last-success.env") {
			val, _ := cmd.Flags().GetString("filter-last-success.env")
			optionalParams.WithFilterLastSuccessEnv(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScannedAssetsMetadata(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-scanned-assets-metadata")

		cmd.Println(cmdutil.FormatJSON(res, "security_monitoring"))
	},
}

func init() {

	ListScannedAssetsMetadataCmd.Flags().String("page-token", "", "Its value must come from the 'links' section of the response of the first request. Do not manually edit it.")

	ListScannedAssetsMetadataCmd.Flags().Int64("page-number", 0, "The page number to be retrieved. It should be equal to or greater than 1.")

	ListScannedAssetsMetadataCmd.Flags().String("filter-asset.type", "", "The type of the scanned asset.")

	ListScannedAssetsMetadataCmd.Flags().String("filter-asset.name", "", "The name of the scanned asset.")

	ListScannedAssetsMetadataCmd.Flags().String("filter-last-success.origin", "", "The origin of last success scan.")

	ListScannedAssetsMetadataCmd.Flags().String("filter-last-success.env", "", "The environment of last success scan.")

	Cmd.AddCommand(ListScannedAssetsMetadataCmd)
}
