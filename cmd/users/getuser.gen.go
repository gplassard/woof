package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserCmd = &cobra.Command{
	Use:   "getuser",
	Short: "Get user details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/users/{user_id}")
		fmt.Println("OperationID: GetUser")
	},
}

func init() {
	Cmd.AddCommand(GetUserCmd)
}
