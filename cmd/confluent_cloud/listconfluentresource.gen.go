package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListConfluentResourceCmd = &cobra.Command{
	Use:   "listconfluentresource",
	Short: "List Confluent Account resources",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integrations/confluent-cloud/accounts/{account_id}/resources")
		fmt.Println("OperationID: ListConfluentResource")
	},
}

func init() {
	Cmd.AddCommand(ListConfluentResourceCmd)
}
