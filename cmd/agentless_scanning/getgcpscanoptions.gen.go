package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetGcpScanOptionsCmd = &cobra.Command{
	Use:   "getgcpscanoptions",
	Short: "Get GCP scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/agentless_scanning/accounts/gcp/{project_id}")
		fmt.Println("OperationID: GetGcpScanOptions")
	},
}

func init() {
	Cmd.AddCommand(GetGcpScanOptionsCmd)
}
