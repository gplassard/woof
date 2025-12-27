package actions_datastores

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDatastoreCmd = &cobra.Command{
	Use:   "deletedatastore [datastore_id]",
	Short: "Delete datastore",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		_, err := api.DeleteDatastore(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletedatastore: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteDatastoreCmd)
}
