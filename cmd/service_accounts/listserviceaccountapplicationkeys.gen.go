package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListServiceAccountApplicationKeysCmd = &cobra.Command{
	Use:   "listserviceaccountapplicationkeys",
	Short: "List application keys for this service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/service_accounts/{service_account_id}/application_keys")
		fmt.Println("OperationID: ListServiceAccountApplicationKeys")
	},
}

func init() {
	Cmd.AddCommand(ListServiceAccountApplicationKeysCmd)
}
