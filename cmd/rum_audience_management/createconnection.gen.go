package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateConnectionCmd = &cobra.Command{
	Use:   "createconnection",
	Short: "Create connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/{entity}/mapping/connection")
		fmt.Println("OperationID: CreateConnection")
	},
}

func init() {
	Cmd.AddCommand(CreateConnectionCmd)
}
