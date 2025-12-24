package software_catalog

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCatalogKindCmd = &cobra.Command{
	Use:   "deletecatalogkind [kind_id]",
	Short: "Delete a single kind",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		_, err := api.DeleteCatalogKind(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecatalogkind: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCatalogKindCmd)
}
