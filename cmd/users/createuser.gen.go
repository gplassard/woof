package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateUserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "Create a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/users")
		fmt.Println("OperationID: CreateUser")
	},
}

func init() {
	Cmd.AddCommand(CreateUserCmd)
}
