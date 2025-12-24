package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAwsScanOptionsCmd = &cobra.Command{
	Use:   "getawsscanoptions",
	Short: "Get AWS scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/aws/{account_id}")
		fmt.Println("OperationID: GetAwsScanOptions")
	},
}

func init() {
	Cmd.AddCommand(GetAwsScanOptionsCmd)
}
