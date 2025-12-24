package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DisableUserCmd = &cobra.Command{
	Use:   "disableuser",
	Short: "Disable a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/users/{user_id}")
		fmt.Println("OperationID: DisableUser")
	},
}

func init() {
	Cmd.AddCommand(DisableUserCmd)
}
