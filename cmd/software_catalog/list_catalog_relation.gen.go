package software_catalog

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogRelationCmd = &cobra.Command{
	Use: "list-catalog-relation",

	Short: "Get a list of entity relations",
	Long: `Get a list of entity relations
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#list-catalog-relation`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListRelationCatalogResponse
		var err error

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCatalogRelation(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-catalog-relation")

		fmt.Println(cmdutil.FormatJSON(res, "catalog_relation"))
	},
}

func init() {

	Cmd.AddCommand(ListCatalogRelationCmd)
}
