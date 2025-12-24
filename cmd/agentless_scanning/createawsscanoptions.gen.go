package agentless_scanning

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAwsScanOptionsCmd = &cobra.Command{
	Use:   "createawsscanoptions",
	Short: "Create AWS scan options",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/agentless_scanning/accounts/aws")
		fmt.Println("OperationID: CreateAwsScanOptions")
	},
}

func init() {
	Cmd.AddCommand(CreateAwsScanOptionsCmd)
}
