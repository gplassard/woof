package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateUserCmd = &cobra.Command{
	Use:   "updateuser",
	Short: "Update a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/users/{user_id}")
		fmt.Println("OperationID: UpdateUser")
	},
}

func init() {
	Cmd.AddCommand(UpdateUserCmd)
}
