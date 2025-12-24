package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDatastoreCmd = &cobra.Command{
	Use:   "updatedatastore",
	Short: "Update datastore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/actions-datastores/{datastore_id}")
		fmt.Println("OperationID: UpdateDatastore")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatastoreCmd)
}
