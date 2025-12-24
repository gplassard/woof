package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetConfluentAccountCmd = &cobra.Command{
	Use:   "getconfluentaccount",
	Short: "Get Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/confluent-cloud/accounts/{account_id}")
		fmt.Println("OperationID: GetConfluentAccount")
	},
}

func init() {
	Cmd.AddCommand(GetConfluentAccountCmd)
}
