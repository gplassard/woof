package okta_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOktaAccountCmd = &cobra.Command{
	Use:   "createoktaaccount",
	Short: "Add Okta account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/okta/accounts")
		fmt.Println("OperationID: CreateOktaAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateOktaAccountCmd)
}
