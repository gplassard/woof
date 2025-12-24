package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteConnectionCmd = &cobra.Command{
	Use:   "deleteconnection",
	Short: "Delete connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/product-analytics/{entity}/mapping/connection/{id}")
		fmt.Println("OperationID: DeleteConnection")
	},
}

func init() {
	Cmd.AddCommand(DeleteConnectionCmd)
}
