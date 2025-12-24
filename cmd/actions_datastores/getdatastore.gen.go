package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDatastoreCmd = &cobra.Command{
	Use:   "getdatastore",
	Short: "Get datastore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions-datastores/{datastore_id}")
		fmt.Println("OperationID: GetDatastore")
	},
}

func init() {
	Cmd.AddCommand(GetDatastoreCmd)
}
