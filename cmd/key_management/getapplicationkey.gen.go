package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetApplicationKeyCmd = &cobra.Command{
	Use:   "getapplicationkey",
	Short: "Get an application key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/application_keys/{app_key_id}")
		fmt.Println("OperationID: GetApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(GetApplicationKeyCmd)
}
