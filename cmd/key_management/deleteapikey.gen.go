package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAPIKeyCmd = &cobra.Command{
	Use:   "deleteapikey",
	Short: "Delete an API key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/api_keys/{api_key_id}")
		fmt.Println("OperationID: DeleteAPIKey")
	},
}

func init() {
	Cmd.AddCommand(DeleteAPIKeyCmd)
}
