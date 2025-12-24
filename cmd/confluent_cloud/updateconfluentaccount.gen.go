package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateConfluentAccountCmd = &cobra.Command{
	Use:   "updateconfluentaccount",
	Short: "Update Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integrations/confluent-cloud/accounts/{account_id}")
		fmt.Println("OperationID: UpdateConfluentAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentAccountCmd)
}
