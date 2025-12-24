package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAwsScanOptionsCmd = &cobra.Command{
	Use:   "updateawsscanoptions",
	Short: "Update AWS scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/agentless_scanning/accounts/aws/{account_id}")
		fmt.Println("OperationID: UpdateAwsScanOptions")
	},
}

func init() {
	Cmd.AddCommand(UpdateAwsScanOptionsCmd)
}
