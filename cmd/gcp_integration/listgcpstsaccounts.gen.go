package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListGCPSTSAccountsCmd = &cobra.Command{
	Use:   "listgcpstsaccounts",
	Short: "List all GCP STS-enabled service accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/gcp/accounts")
		fmt.Println("OperationID: ListGCPSTSAccounts")
	},
}

func init() {
	Cmd.AddCommand(ListGCPSTSAccountsCmd)
}
