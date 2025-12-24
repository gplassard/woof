package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "updateserviceaccountapplicationkey",
	Short: "Edit an application key for this service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}")
		fmt.Println("OperationID: UpdateServiceAccountApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(UpdateServiceAccountApplicationKeyCmd)
}
