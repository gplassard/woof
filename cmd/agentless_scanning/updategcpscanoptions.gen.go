package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateGcpScanOptionsCmd = &cobra.Command{
	Use:   "updategcpscanoptions",
	Short: "Update GCP scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/agentless_scanning/accounts/gcp/{project_id}")
		fmt.Println("OperationID: UpdateGcpScanOptions")
	},
}

func init() {
	Cmd.AddCommand(UpdateGcpScanOptionsCmd)
}
