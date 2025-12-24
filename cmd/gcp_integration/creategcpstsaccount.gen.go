package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateGCPSTSAccountCmd = &cobra.Command{
	Use:   "creategcpstsaccount",
	Short: "Create a new entry for your service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/gcp/accounts")
		fmt.Println("OperationID: CreateGCPSTSAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateGCPSTSAccountCmd)
}
