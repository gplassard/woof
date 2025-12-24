package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRoleCmd = &cobra.Command{
	Use:   "getrole",
	Short: "Get a role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/roles/{role_id}")
		fmt.Println("OperationID: GetRole")
	},
}

func init() {
	Cmd.AddCommand(GetRoleCmd)
}
