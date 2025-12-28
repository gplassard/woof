package actions_datastores

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDatastoreItemCmd = &cobra.Command{
	Use:   "delete-datastore-item [datastore_id]",
	Short: "Delete datastore item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.DeleteDatastoreItem(client.NewContext(apiKey, appKey, site), args[0], datadogV2.DeleteAppsDatastoreItemRequest{})
		if err != nil {
			log.Fatalf("failed to delete-datastore-item: %v", err)
		}

		cmdutil.PrintJSON(res, "items")
	},
}

func init() {
	Cmd.AddCommand(DeleteDatastoreItemCmd)
}
