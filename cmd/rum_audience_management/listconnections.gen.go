package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListConnectionsCmd = &cobra.Command{
	Use:   "listconnections",
	Short: "List connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/product-analytics/{entity}/mapping/connections")
		fmt.Println("OperationID: ListConnections")
	},
}

func init() {
	Cmd.AddCommand(ListConnectionsCmd)
}
