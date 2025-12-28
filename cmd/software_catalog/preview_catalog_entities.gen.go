package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var PreviewCatalogEntitiesCmd = &cobra.Command{
	Use:   "preview-catalog-entities",
	Short: "Preview catalog entities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.PreviewCatalogEntities(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to preview-catalog-entities: %v", err)
		}

		cmdutil.PrintJSON(res, "entity")
	},
}

func init() {
	Cmd.AddCommand(PreviewCatalogEntitiesCmd)
}
