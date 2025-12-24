package fastly_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateFastlyServiceCmd = &cobra.Command{
	Use:   "createfastlyservice",
	Short: "Add Fastly service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/fastly/accounts/{account_id}/services")
		fmt.Println("OperationID: CreateFastlyService")
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyServiceCmd)
}
