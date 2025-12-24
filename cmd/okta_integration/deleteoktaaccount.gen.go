package okta_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOktaAccountCmd = &cobra.Command{
	Use:   "deleteoktaaccount",
	Short: "Delete Okta account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/okta/accounts/{account_id}")
		fmt.Println("OperationID: DeleteOktaAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteOktaAccountCmd)
}
