package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCatalogEntityCmd = &cobra.Command{
	Use:   "listcatalogentity",
	Short: "Get a list of entities",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.ListCatalogEntity(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listcatalogentity: %v", err)
		}

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogEntityCmd)
}
