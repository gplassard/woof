package okta_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOktaAccountCmd = &cobra.Command{
	Use:   "updateoktaaccount",
	Short: "Update Okta account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/okta/accounts/{account_id}")
		fmt.Println("OperationID: UpdateOktaAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateOktaAccountCmd)
}
