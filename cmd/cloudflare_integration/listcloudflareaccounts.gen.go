package cloudflare_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCloudflareAccountsCmd = &cobra.Command{
	Use:   "listcloudflareaccounts",
	Short: "List Cloudflare accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/cloudflare/accounts")
		fmt.Println("OperationID: ListCloudflareAccounts")
	},
}

func init() {
	Cmd.AddCommand(ListCloudflareAccountsCmd)
}
