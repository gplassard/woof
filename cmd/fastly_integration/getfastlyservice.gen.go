package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFastlyServiceCmd = &cobra.Command{
	Use:   "getfastlyservice",
	Short: "Get Fastly service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/fastly/accounts/{account_id}/services/{service_id}")
		fmt.Println("OperationID: GetFastlyService")
	},
}

func init() {
	Cmd.AddCommand(GetFastlyServiceCmd)
}
