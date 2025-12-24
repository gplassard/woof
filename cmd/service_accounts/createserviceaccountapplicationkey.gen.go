package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "createserviceaccountapplicationkey",
	Short: "Create an application key for this service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/service_accounts/{service_account_id}/application_keys")
		fmt.Println("OperationID: CreateServiceAccountApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(CreateServiceAccountApplicationKeyCmd)
}
