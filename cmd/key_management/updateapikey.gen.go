package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAPIKeyCmd = &cobra.Command{
	Use:   "updateapikey",
	Short: "Edit an API key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/api_keys/{api_key_id}")
		fmt.Println("OperationID: UpdateAPIKey")
	},
}

func init() {
	Cmd.AddCommand(UpdateAPIKeyCmd)
}
