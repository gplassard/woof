package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMappingCmd = &cobra.Command{
	Use:   "getmapping",
	Short: "Get mapping",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/product-analytics/{entity}/mapping")
		fmt.Println("OperationID: GetMapping")
	},
}

func init() {
	Cmd.AddCommand(GetMappingCmd)
}
