package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAzureScanOptionsCmd = &cobra.Command{
	Use:   "deleteazurescanoptions",
	Short: "Delete Azure scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/agentless_scanning/accounts/azure/{subscription_id}")
		fmt.Println("OperationID: DeleteAzureScanOptions")
	},
}

func init() {
	Cmd.AddCommand(DeleteAzureScanOptionsCmd)
}
