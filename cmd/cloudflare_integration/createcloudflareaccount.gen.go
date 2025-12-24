package cloudflare_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCloudflareAccountCmd = &cobra.Command{
	Use:   "createcloudflareaccount",
	Short: "Add Cloudflare account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/cloudflare/accounts")
		fmt.Println("OperationID: CreateCloudflareAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateCloudflareAccountCmd)
}
