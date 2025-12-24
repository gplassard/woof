package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAPIKeysCmd = &cobra.Command{
	Use:   "listapikeys",
	Short: "Get all API keys",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/api_keys")
		fmt.Println("OperationID: ListAPIKeys")
	},
}

func init() {
	Cmd.AddCommand(ListAPIKeysCmd)
}
