package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDatastoreItemCmd = &cobra.Command{
	Use:   "deletedatastoreitem",
	Short: "Delete datastore item",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/actions-datastores/{datastore_id}/items")
		fmt.Println("OperationID: DeleteDatastoreItem")
	},
}

func init() {
	Cmd.AddCommand(DeleteDatastoreItemCmd)
}
