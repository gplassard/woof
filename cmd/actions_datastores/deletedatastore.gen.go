package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDatastoreCmd = &cobra.Command{
	Use:   "deletedatastore",
	Short: "Delete datastore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/actions-datastores/{datastore_id}")
		fmt.Println("OperationID: DeleteDatastore")
	},
}

func init() {
	Cmd.AddCommand(DeleteDatastoreCmd)
}
