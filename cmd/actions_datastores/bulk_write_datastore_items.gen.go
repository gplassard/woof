package actions_datastores

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var BulkWriteDatastoreItemsCmd = &cobra.Command{
	Use:   "bulk_write_datastore_items [datastore_id]",
	Short: "Bulk write datastore items",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.BulkWriteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BulkPutAppsDatastoreItemsRequest{})
		if err != nil {
			log.Fatalf("failed to bulk_write_datastore_items: %v", err)
		}

		cmdutil.PrintJSON(res, "actions_datastores")
	},
}

func init() {
	Cmd.AddCommand(BulkWriteDatastoreItemsCmd)
}
