package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAPIKeyCmd = &cobra.Command{
	Use:   "getapikey",
	Short: "Get API key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/api_keys/{api_key_id}")
		fmt.Println("OperationID: GetAPIKey")
	},
}

func init() {
	Cmd.AddCommand(GetAPIKeyCmd)
}
