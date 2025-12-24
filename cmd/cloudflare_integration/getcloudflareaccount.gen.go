package cloudflare_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCloudflareAccountCmd = &cobra.Command{
	Use:   "getcloudflareaccount",
	Short: "Get Cloudflare account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/cloudflare/accounts/{account_id}")
		fmt.Println("OperationID: GetCloudflareAccount")
	},
}

func init() {
	Cmd.AddCommand(GetCloudflareAccountCmd)
}
