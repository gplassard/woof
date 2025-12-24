package actions_datastores

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var BulkWriteDatastoreItemsCmd = &cobra.Command{
	Use:   "bulkwritedatastoreitems [datastore_id]",
	Short: "Bulk write datastore items",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewActionsDatastoresApi(client.NewAPIClient())
		res, _, err := api.BulkWriteDatastoreItems(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BulkPutAppsDatastoreItemsRequest{})
		if err != nil {
			log.Fatalf("failed to bulkwritedatastoreitems: %v", err)
		}

		cmdutil.PrintJSON(res, "actions_datastores")
	},
}

func init() {
	Cmd.AddCommand(BulkWriteDatastoreItemsCmd)
}
