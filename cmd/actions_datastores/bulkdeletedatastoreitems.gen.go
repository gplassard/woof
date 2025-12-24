package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var BulkDeleteDatastoreItemsCmd = &cobra.Command{
	Use:   "bulkdeletedatastoreitems",
	Short: "Bulk delete datastore items",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/actions-datastores/{datastore_id}/items/bulk")
		fmt.Println("OperationID: BulkDeleteDatastoreItems")
	},
}

func init() {
	Cmd.AddCommand(BulkDeleteDatastoreItemsCmd)
}
