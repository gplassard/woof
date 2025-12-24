package api_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAPIsCmd = &cobra.Command{
	Use:   "listapis",
	Short: "List APIs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apicatalog/api")
		fmt.Println("OperationID: ListAPIs")
	},
}

func init() {
	Cmd.AddCommand(ListAPIsCmd)
}
