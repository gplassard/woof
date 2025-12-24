package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteFastlyServiceCmd = &cobra.Command{
	Use:   "deletefastlyservice",
	Short: "Delete Fastly service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/fastly/accounts/{account_id}/services/{service_id}")
		fmt.Println("OperationID: DeleteFastlyService")
	},
}

func init() {
	Cmd.AddCommand(DeleteFastlyServiceCmd)
}
