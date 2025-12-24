package okta_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetOktaAccountCmd = &cobra.Command{
	Use:   "getoktaaccount",
	Short: "Get Okta account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/okta/accounts/{account_id}")
		fmt.Println("OperationID: GetOktaAccount")
	},
}

func init() {
	Cmd.AddCommand(GetOktaAccountCmd)
}
