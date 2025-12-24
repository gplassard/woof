package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetConfluentResourceCmd = &cobra.Command{
	Use:   "getconfluentresource",
	Short: "Get resource from Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/confluent-cloud/accounts/{account_id}/resources/{resource_id}")
		fmt.Println("OperationID: GetConfluentResource")
	},
}

func init() {
	Cmd.AddCommand(GetConfluentResourceCmd)
}
