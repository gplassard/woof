package security_monitoring

import (
	"fmt"
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

		optionalParams := datadogV2.NewGetSBOMOptionalParameters()

		if cmd.Flags().Changed("filter-repo-digest") {
			val, _ := cmd.Flags().GetString("filter-repo-digest")
			optionalParams.WithFilterRepoDigest(val)
		}

		if cmd.Flags().Changed("ext:format") {
			val, _ := cmd.Flags().GetString("ext:format")
			optionalParams.WithExtFormat(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSBOM(client.NewContext(apiKey, appKey, site), datadogV2.AssetType(args[0]), args[1], *optionalParams)
		cmdutil.HandleError(err, "failed to get-sbom")

		fmt.Println(cmdutil.FormatJSON(res, "sboms"))
	},
}

func init() {

	GetSBOMCmd.Flags().String("filter-repo-digest", "", "The container image 'repo_digest' for the SBOM request. When the requested asset type is 'Image', this filter is mandatory.")

	GetSBOMCmd.Flags().String("ext:format", "", "The standard of the SBOM.")

	Cmd.AddCommand(GetSBOMCmd)
}
