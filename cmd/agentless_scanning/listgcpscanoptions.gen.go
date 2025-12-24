package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListGcpScanOptionsCmd = &cobra.Command{
	Use:   "listgcpscanoptions",
	Short: "List GCP scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/gcp")
		fmt.Println("OperationID: ListGcpScanOptions")
	},
}

func init() {
	Cmd.AddCommand(ListGcpScanOptionsCmd)
}
