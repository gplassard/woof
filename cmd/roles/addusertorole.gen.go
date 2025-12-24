package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddUserToRoleCmd = &cobra.Command{
	Use:   "addusertorole",
	Short: "Add a user to a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/roles/{role_id}/users")
		fmt.Println("OperationID: AddUserToRole")
	},
}

func init() {
	Cmd.AddCommand(AddUserToRoleCmd)
}
