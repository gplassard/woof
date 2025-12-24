package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAzureScanOptionsCmd = &cobra.Command{
	Use:   "listazurescanoptions",
	Short: "List Azure scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/azure")
		fmt.Println("OperationID: ListAzureScanOptions")
	},
}

func init() {
	Cmd.AddCommand(ListAzureScanOptionsCmd)
}
