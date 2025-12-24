package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteConfluentAccountCmd = &cobra.Command{
	Use:   "deleteconfluentaccount",
	Short: "Delete Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/confluent-cloud/accounts/{account_id}")
		fmt.Println("OperationID: DeleteConfluentAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteConfluentAccountCmd)
}
