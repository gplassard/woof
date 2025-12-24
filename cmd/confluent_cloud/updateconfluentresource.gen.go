package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateConfluentResourceCmd = &cobra.Command{
	Use:   "updateconfluentresource",
	Short: "Update resource in Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/confluent-cloud/accounts/{account_id}/resources/{resource_id}")
		fmt.Println("OperationID: UpdateConfluentResource")
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentResourceCmd)
}
