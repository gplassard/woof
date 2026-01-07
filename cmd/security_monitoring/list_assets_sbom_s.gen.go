package security_monitoring

import (
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

		optionalParams := datadogV2.NewListAssetsSBOMsOptionalParameters()

		if cmd.Flags().Changed("page-token") {
			val, _ := cmd.Flags().GetString("page-token")
			optionalParams.WithPageToken(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("filter-asset-type") {
			val, _ := cmd.Flags().GetString("filter-asset-type")
			optionalParams.WithFilterAssetType(val)
		}

		if cmd.Flags().Changed("filter-asset-name") {
			val, _ := cmd.Flags().GetString("filter-asset-name")
			optionalParams.WithFilterAssetName(val)
		}

		if cmd.Flags().Changed("filter-package-name") {
			val, _ := cmd.Flags().GetString("filter-package-name")
			optionalParams.WithFilterPackageName(val)
		}

		if cmd.Flags().Changed("filter-package-version") {
			val, _ := cmd.Flags().GetString("filter-package-version")
			optionalParams.WithFilterPackageVersion(val)
		}

		if cmd.Flags().Changed("filter-license-name") {
			val, _ := cmd.Flags().GetString("filter-license-name")
			optionalParams.WithFilterLicenseName(val)
		}

		if cmd.Flags().Changed("filter-license-type") {
			val, _ := cmd.Flags().GetString("filter-license-type")
			optionalParams.WithFilterLicenseType(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAssetsSBOMs(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-assets-sbom-s")

		cmd.Println(cmdutil.FormatJSON(res, "sboms"))
	},
}

func init() {

	ListAssetsSBOMsCmd.Flags().String("page-token", "", "Its value must come from the 'links' section of the response of the first request. Do not manually edit it.")

	ListAssetsSBOMsCmd.Flags().Int64("page-number", 0, "The page number to be retrieved. It should be equal to or greater than 1.")

	ListAssetsSBOMsCmd.Flags().String("filter-asset-type", "", "The type of the assets for the SBOM request.")

	ListAssetsSBOMsCmd.Flags().String("filter-asset-name", "", "The name of the asset for the SBOM request.")

	ListAssetsSBOMsCmd.Flags().String("filter-package-name", "", "The name of the component that is a dependency of an asset.")

	ListAssetsSBOMsCmd.Flags().String("filter-package-version", "", "The version of the component that is a dependency of an asset.")

	ListAssetsSBOMsCmd.Flags().String("filter-license-name", "", "The software license name of the component that is a dependency of an asset.")

	ListAssetsSBOMsCmd.Flags().String("filter-license-type", "", "The software license type of the component that is a dependency of an asset.")

	Cmd.AddCommand(ListAssetsSBOMsCmd)
}
