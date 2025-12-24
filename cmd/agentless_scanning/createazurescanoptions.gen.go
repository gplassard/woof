package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAzureScanOptionsCmd = &cobra.Command{
	Use:   "createazurescanoptions",
	Short: "Create Azure scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/agentless_scanning/accounts/azure")
		fmt.Println("OperationID: CreateAzureScanOptions")
	},
}

func init() {
	Cmd.AddCommand(CreateAzureScanOptionsCmd)
}
