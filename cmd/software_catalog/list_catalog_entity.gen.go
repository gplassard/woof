package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogEntityCmd = &cobra.Command{
	Use: "list-catalog-entity",

	Short: "Get a list of entities",
	Long: `Get a list of entities
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#list-catalog-entity`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListEntityCatalogResponse
		var err error

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err = api.ListCatalogEntity(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-catalog-entity")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	Cmd.AddCommand(ListCatalogEntityCmd)
}
