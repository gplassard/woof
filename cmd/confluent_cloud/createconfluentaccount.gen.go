package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateConfluentAccountCmd = &cobra.Command{
	Use:   "createconfluentaccount",
	Short: "Add Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/confluent-cloud/accounts")
		fmt.Println("OperationID: CreateConfluentAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentAccountCmd)
}
