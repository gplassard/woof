package service_accounts

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateServiceAccountCmd = &cobra.Command{
	Use:   "createserviceaccount",
	Short: "Create a service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/service_accounts")
		fmt.Println("OperationID: CreateServiceAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateServiceAccountCmd)
}
