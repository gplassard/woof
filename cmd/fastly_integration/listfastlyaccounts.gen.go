package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFastlyAccountsCmd = &cobra.Command{
	Use:   "listfastlyaccounts",
	Short: "List Fastly accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/fastly/accounts")
		fmt.Println("OperationID: ListFastlyAccounts")
	},
}

func init() {
	Cmd.AddCommand(ListFastlyAccountsCmd)
}
