package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAWSAccountsCmd = &cobra.Command{
	Use:   "listawsaccounts",
	Short: "List all AWS integrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/accounts")
		fmt.Println("OperationID: ListAWSAccounts")
	},
}

func init() {
	Cmd.AddCommand(ListAWSAccountsCmd)
}
