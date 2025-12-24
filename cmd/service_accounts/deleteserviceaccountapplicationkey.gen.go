package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "deleteserviceaccountapplicationkey",
	Short: "Delete an application key for this service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}")
		fmt.Println("OperationID: DeleteServiceAccountApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(DeleteServiceAccountApplicationKeyCmd)
}
