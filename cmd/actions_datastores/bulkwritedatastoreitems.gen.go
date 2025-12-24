package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var BulkWriteDatastoreItemsCmd = &cobra.Command{
	Use:   "bulkwritedatastoreitems",
	Short: "Bulk write datastore items",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/actions-datastores/{datastore_id}/items/bulk")
		fmt.Println("OperationID: BulkWriteDatastoreItems")
	},
}

func init() {
	Cmd.AddCommand(BulkWriteDatastoreItemsCmd)
}
