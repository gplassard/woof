package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteConfluentResourceCmd = &cobra.Command{
	Use:   "deleteconfluentresource",
	Short: "Delete resource from Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integrations/confluent-cloud/accounts/{account_id}/resources/{resource_id}")
		fmt.Println("OperationID: DeleteConfluentResource")
	},
}

func init() {
	Cmd.AddCommand(DeleteConfluentResourceCmd)
}
