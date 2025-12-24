package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAzureScanOptionsCmd = &cobra.Command{
	Use:   "getazurescanoptions",
	Short: "Get Azure scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/azure/{subscription_id}")
		fmt.Println("OperationID: GetAzureScanOptions")
	},
}

func init() {
	Cmd.AddCommand(GetAzureScanOptionsCmd)
}
