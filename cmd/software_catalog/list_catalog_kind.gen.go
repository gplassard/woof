package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCatalogKindCmd = &cobra.Command{
	Use:   "list-catalog-kind",
	Short: "Get a list of entity kinds",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.ListCatalogKind(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-catalog-kind: %v", err)
		}

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogKindCmd)
}
