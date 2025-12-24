package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var QueryUsersCmd = &cobra.Command{
	Use:   "queryusers",
	Short: "Query users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/users/query")
		fmt.Println("OperationID: QueryUsers")
	},
}

func init() {
	Cmd.AddCommand(QueryUsersCmd)
}
