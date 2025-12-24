package actions_datastores

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDatastoresCmd = &cobra.Command{
	Use:   "listdatastores",
	Short: "List datastores",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/actions-datastores")
		fmt.Println("OperationID: ListDatastores")
	},
}

func init() {
	Cmd.AddCommand(ListDatastoresCmd)
}
