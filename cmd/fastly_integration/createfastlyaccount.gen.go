package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateFastlyAccountCmd = &cobra.Command{
	Use:   "createfastlyaccount",
	Short: "Add Fastly account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/fastly/accounts")
		fmt.Println("OperationID: CreateFastlyAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyAccountCmd)
}
