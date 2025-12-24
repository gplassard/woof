package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListApplicationKeysCmd = &cobra.Command{
	Use:   "listapplicationkeys",
	Short: "Get all application keys",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/application_keys")
		fmt.Println("OperationID: ListApplicationKeys")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationKeysCmd)
}
