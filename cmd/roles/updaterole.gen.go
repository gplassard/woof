package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRoleCmd = &cobra.Command{
	Use:   "updaterole",
	Short: "Update a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/roles/{role_id}")
		fmt.Println("OperationID: UpdateRole")
	},
}

func init() {
	Cmd.AddCommand(UpdateRoleCmd)
}
