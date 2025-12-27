package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpsertCatalogEntityCmd = &cobra.Command{
	Use:   "upsertcatalogentity",
	Short: "Create or update entities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.UpsertCatalogEntity(client.NewContext(apiKey, appKey, site), datadogV2.UpsertCatalogEntityRequest{})
		if err != nil {
			log.Fatalf("failed to upsertcatalogentity: %v", err)
		}

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogEntityCmd)
}
