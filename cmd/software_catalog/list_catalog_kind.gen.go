package software_catalog

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogKindCmd = &cobra.Command{
	Use: "list-catalog-kind",

	Short: "Get a list of entity kinds",
	Long: `Get a list of entity kinds
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#list-catalog-kind`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListKindCatalogResponse
		var err error

		optionalParams := datadogV2.NewListCatalogKindOptionalParameters()

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("filter-id") {
			val, _ := cmd.Flags().GetString("filter-id")
			optionalParams.WithFilterId(val)
		}

		if cmd.Flags().Changed("filter-name") {
			val, _ := cmd.Flags().GetString("filter-name")
			optionalParams.WithFilterName(val)
		}

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCatalogKind(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-catalog-kind")

		fmt.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	ListCatalogKindCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListCatalogKindCmd.Flags().Int64("page-limit", 0, "Maximum number of kinds in the response.")

	ListCatalogKindCmd.Flags().String("filter-id", "", "Filter entities by UUID.")

	ListCatalogKindCmd.Flags().String("filter-name", "", "Filter entities by name.")

	Cmd.AddCommand(ListCatalogKindCmd)
}
