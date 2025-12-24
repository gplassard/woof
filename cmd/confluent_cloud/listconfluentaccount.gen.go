package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListConfluentAccountCmd = &cobra.Command{
	Use:   "listconfluentaccount",
	Short: "List Confluent accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/confluent-cloud/accounts")
		fmt.Println("OperationID: ListConfluentAccount")
	},
}

func init() {
	Cmd.AddCommand(ListConfluentAccountCmd)
}
