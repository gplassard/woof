package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAwsScanOptionsCmd = &cobra.Command{
	Use:   "listawsscanoptions",
	Short: "List AWS scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/aws")
		fmt.Println("OperationID: ListAwsScanOptions")
	},
}

func init() {
	Cmd.AddCommand(ListAwsScanOptionsCmd)
}
