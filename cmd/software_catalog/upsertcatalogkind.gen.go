package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpsertCatalogKindCmd = &cobra.Command{
	Use:   "upsertcatalogkind",
	Short: "Create or update kinds",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.UpsertCatalogKind(client.NewContext(apiKey, appKey, site), datadogV2.UpsertCatalogKindRequest{})
		if err != nil {
			log.Fatalf("failed to upsertcatalogkind: %v", err)
		}

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogKindCmd)
}
