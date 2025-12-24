package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateGcpScanOptionsCmd = &cobra.Command{
	Use:   "creategcpscanoptions",
	Short: "Create GCP scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/agentless_scanning/accounts/gcp")
		fmt.Println("OperationID: CreateGcpScanOptions")
	},
}

func init() {
	Cmd.AddCommand(CreateGcpScanOptionsCmd)
}
