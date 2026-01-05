package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCatalogEntityCmd = &cobra.Command{
	Use: "delete-catalog-entity [entity_id]",

	Short: "Delete a single entity",
	Long: `Delete a single entity
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#delete-catalog-entity`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		_, err = api.DeleteCatalogEntity(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-catalog-entity")

	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogEntityCmd)
}
