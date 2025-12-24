package cloudflare_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCloudflareAccountCmd = &cobra.Command{
	Use:   "deletecloudflareaccount",
	Short: "Delete Cloudflare account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/cloudflare/accounts/{account_id}")
		fmt.Println("OperationID: DeleteCloudflareAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteCloudflareAccountCmd)
}
