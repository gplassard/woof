package software_catalog

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var PreviewCatalogEntitiesCmd = &cobra.Command{
	Use: "preview-catalog-entities",

	Short: "Preview catalog entities",
	Long: `Preview catalog entities
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#preview-catalog-entities`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EntityResponseArray
		var err error

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.PreviewCatalogEntities(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to preview-catalog-entities")

		fmt.Println(cmdutil.FormatJSON(res, "entity"))
	},
}

func init() {

	Cmd.AddCommand(PreviewCatalogEntitiesCmd)
}
