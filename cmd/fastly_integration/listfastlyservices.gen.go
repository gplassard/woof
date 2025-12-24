package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFastlyServicesCmd = &cobra.Command{
	Use:   "listfastlyservices",
	Short: "List Fastly services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/fastly/accounts/{account_id}/services")
		fmt.Println("OperationID: ListFastlyServices")
	},
}

func init() {
	Cmd.AddCommand(ListFastlyServicesCmd)
}
