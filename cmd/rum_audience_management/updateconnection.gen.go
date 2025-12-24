package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateConnectionCmd = &cobra.Command{
	Use:   "updateconnection",
	Short: "Update connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/product-analytics/{entity}/mapping/connection")
		fmt.Println("OperationID: UpdateConnection")
	},
}

func init() {
	Cmd.AddCommand(UpdateConnectionCmd)
}
