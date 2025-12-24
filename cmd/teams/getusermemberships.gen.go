package teams

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserMembershipsCmd = &cobra.Command{
	Use:   "getusermemberships",
	Short: "Get user memberships",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/users/{user_uuid}/memberships")
		fmt.Println("OperationID: GetUserMemberships")
	},
}

func init() {
	Cmd.AddCommand(GetUserMembershipsCmd)
}
