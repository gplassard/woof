package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateFastlyServiceCmd = &cobra.Command{
	Use:   "updatefastlyservice",
	Short: "Update Fastly service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/fastly/accounts/{account_id}/services/{service_id}")
		fmt.Println("OperationID: UpdateFastlyService")
	},
}

func init() {
	Cmd.AddCommand(UpdateFastlyServiceCmd)
}
