package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CloneRoleCmd = &cobra.Command{
	Use:   "clonerole",
	Short: "Create a new role by cloning an existing role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/roles/{role_id}/clone")
		fmt.Println("OperationID: CloneRole")
	},
}

func init() {
	Cmd.AddCommand(CloneRoleCmd)
}
