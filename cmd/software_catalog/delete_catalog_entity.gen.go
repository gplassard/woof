package software_catalog

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCatalogEntityCmd = &cobra.Command{
	Use:   "delete-catalog-entity [entity_id]",
	Short: "Delete a single entity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		_, err := api.DeleteCatalogEntity(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-catalog-entity: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogEntityCmd)
}
