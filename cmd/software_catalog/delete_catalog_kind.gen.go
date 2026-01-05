package software_catalog

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCatalogKindCmd = &cobra.Command{
	Use: "delete-catalog-kind [kind_id]",

	Short: "Delete a single kind",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		_, err := api.DeleteCatalogKind(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-catalog-kind")

	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogKindCmd)
}
