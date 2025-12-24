package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFastlyAccountCmd = &cobra.Command{
	Use:   "getfastlyaccount",
	Short: "Get Fastly account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/fastly/accounts/{account_id}")
		fmt.Println("OperationID: GetFastlyAccount")
	},
}

func init() {
	Cmd.AddCommand(GetFastlyAccountCmd)
}
