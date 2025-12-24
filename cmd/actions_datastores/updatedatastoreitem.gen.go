package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDatastoreItemCmd = &cobra.Command{
	Use:   "updatedatastoreitem",
	Short: "Update datastore item",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/actions-datastores/{datastore_id}/items")
		fmt.Println("OperationID: UpdateDatastoreItem")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreItemCmd)
}
