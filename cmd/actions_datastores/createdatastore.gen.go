package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDatastoreCmd = &cobra.Command{
	Use:   "createdatastore",
	Short: "Create datastore",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/actions-datastores")
		fmt.Println("OperationID: CreateDatastore")
	},
}

func init() {
	Cmd.AddCommand(CreateDatastoreCmd)
}
