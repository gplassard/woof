package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "deletecurrentuserapplicationkey",
	Short: "Delete an application key owned by current user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/current_user/application_keys/{app_key_id}")
		fmt.Println("OperationID: DeleteCurrentUserApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(DeleteCurrentUserApplicationKeyCmd)
}
