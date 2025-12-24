package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteFastlyAccountCmd = &cobra.Command{
	Use:   "deletefastlyaccount",
	Short: "Delete Fastly account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/fastly/accounts/{account_id}")
		fmt.Println("OperationID: DeleteFastlyAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteFastlyAccountCmd)
}
