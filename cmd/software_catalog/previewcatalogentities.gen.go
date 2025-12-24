package software_catalog

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var PreviewCatalogEntitiesCmd = &cobra.Command{
	Use:   "previewcatalogentities",
	Short: "Preview catalog entities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.PreviewCatalogEntities(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to previewcatalogentities: %v", err)
		}

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(PreviewCatalogEntitiesCmd)
}
