package software_catalog

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var PreviewCatalogEntitiesCmd = &cobra.Command{
	Use: "preview-catalog-entities",

	Short: "Preview catalog entities",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.PreviewCatalogEntities(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to preview-catalog-entities")

		cmd.Println(cmdutil.FormatJSON(res, "entity"))
	},
}

func init() {
	Cmd.AddCommand(PreviewCatalogEntitiesCmd)
}
