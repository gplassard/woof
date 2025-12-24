package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAzureScanOptionsCmd = &cobra.Command{
	Use:   "updateazurescanoptions",
	Short: "Update Azure scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/agentless_scanning/accounts/azure/{subscription_id}")
		fmt.Println("OperationID: UpdateAzureScanOptions")
	},
}

func init() {
	Cmd.AddCommand(UpdateAzureScanOptionsCmd)
}
