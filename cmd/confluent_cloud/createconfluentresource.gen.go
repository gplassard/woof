package confluent_cloud

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateConfluentResourceCmd = &cobra.Command{
	Use:   "createconfluentresource",
	Short: "Add resource to Confluent account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integrations/confluent-cloud/accounts/{account_id}/resources")
		fmt.Println("OperationID: CreateConfluentResource")
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentResourceCmd)
}
