package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAwsScanOptionsCmd = &cobra.Command{
	Use:   "deleteawsscanoptions",
	Short: "Delete AWS scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/agentless_scanning/accounts/aws/{account_id}")
		fmt.Println("OperationID: DeleteAwsScanOptions")
	},
}

func init() {
	Cmd.AddCommand(DeleteAwsScanOptionsCmd)
}
