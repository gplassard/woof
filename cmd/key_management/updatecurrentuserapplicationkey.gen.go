package key_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "updatecurrentuserapplicationkey",
	Short: "Edit an application key owned by current user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/current_user/application_keys/{app_key_id}")
		fmt.Println("OperationID: UpdateCurrentUserApplicationKey")
	},
}

func init() {
	Cmd.AddCommand(UpdateCurrentUserApplicationKeyCmd)
}
