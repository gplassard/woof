package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertCatalogEntityCmd = &cobra.Command{
	Use: "upsert-catalog-entity",

	Short: "Create or update entities",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.UpsertCatalogEntity(client.NewContext(apiKey, appKey, site), datadogV2.UpsertCatalogEntityRequest{})
		cmdutil.HandleError(err, "failed to upsert-catalog-entity")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogEntityCmd)
}
