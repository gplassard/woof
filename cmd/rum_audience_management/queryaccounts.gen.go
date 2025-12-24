package rum_audience_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var QueryAccountsCmd = &cobra.Command{
	Use:   "queryaccounts",
	Short: "Query accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/product-analytics/accounts/query")
		fmt.Println("OperationID: QueryAccounts")
	},
}

func init() {
	Cmd.AddCommand(QueryAccountsCmd)
}
