package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDatastoreItemsCmd = &cobra.Command{
	Use:   "listdatastoreitems",
	Short: "List datastore items",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions-datastores/{datastore_id}/items")
		fmt.Println("OperationID: ListDatastoreItems")
	},
}

func init() {
	Cmd.AddCommand(ListDatastoreItemsCmd)
}
