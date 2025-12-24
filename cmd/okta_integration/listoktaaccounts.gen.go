package okta_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListOktaAccountsCmd = &cobra.Command{
	Use:   "listoktaaccounts",
	Short: "List Okta accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/okta/accounts")
		fmt.Println("OperationID: ListOktaAccounts")
	},
}

func init() {
	Cmd.AddCommand(ListOktaAccountsCmd)
}
