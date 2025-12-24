package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetServiceAccountApplicationKeyCmd = &cobra.Command{
	Use:   "getserviceaccountapplicationkey",
	Short: "Get one application key for this service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/service_accounts/{service_account_id}/application_keys/{app_key_id}")
		fmt.Println("OperationID: GetServiceAccountApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(GetServiceAccountApplicationKeyCmd)
}
