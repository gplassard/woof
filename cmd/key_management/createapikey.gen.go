package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAPIKeyCmd = &cobra.Command{
	Use:   "createapikey",
	Short: "Create an API key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/api_keys")
		fmt.Println("OperationID: CreateAPIKey")
	},
}

func init() {
	Cmd.AddCommand(CreateAPIKeyCmd)
}
