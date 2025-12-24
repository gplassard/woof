package roles

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateRoleCmd = &cobra.Command{
	Use:   "createrole",
	Short: "Create role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/roles")
		fmt.Println("OperationID: CreateRole")
	},
}

func init() {
	Cmd.AddCommand(CreateRoleCmd)
}
