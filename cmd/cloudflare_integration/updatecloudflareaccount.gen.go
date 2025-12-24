package cloudflare_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCloudflareAccountCmd = &cobra.Command{
	Use:   "updatecloudflareaccount",
	Short: "Update Cloudflare account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/cloudflare/accounts/{account_id}")
		fmt.Println("OperationID: UpdateCloudflareAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateCloudflareAccountCmd)
}
