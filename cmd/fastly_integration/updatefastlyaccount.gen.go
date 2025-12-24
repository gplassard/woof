package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateFastlyAccountCmd = &cobra.Command{
	Use:   "updatefastlyaccount",
	Short: "Update Fastly account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/fastly/accounts/{account_id}")
		fmt.Println("OperationID: UpdateFastlyAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyAccountCmd)
}
