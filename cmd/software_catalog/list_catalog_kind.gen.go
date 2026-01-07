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

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCatalogKind(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-catalog-kind")

		fmt.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	Cmd.AddCommand(ListCatalogKindCmd)
}
