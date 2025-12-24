package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetInvitationCmd = &cobra.Command{
	Use:   "getinvitation",
	Short: "Get a user invitation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/user_invitations/{user_invitation_uuid}")
		fmt.Println("OperationID: GetInvitation")
	},
}

func init() {
	Cmd.AddCommand(GetInvitationCmd)
}
