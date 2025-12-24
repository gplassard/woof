package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteGcpScanOptionsCmd = &cobra.Command{
	Use:   "deletegcpscanoptions",
	Short: "Delete GCP scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/agentless_scanning/accounts/gcp/{project_id}")
		fmt.Println("OperationID: DeleteGcpScanOptions")
	},
}

func init() {
	Cmd.AddCommand(DeleteGcpScanOptionsCmd)
}
