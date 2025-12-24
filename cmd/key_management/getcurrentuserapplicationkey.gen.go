package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "getcurrentuserapplicationkey",
	Short: "Get one application key owned by current user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/current_user/application_keys/{app_key_id}")
		fmt.Println("OperationID: GetCurrentUserApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(GetCurrentUserApplicationKeyCmd)
}
