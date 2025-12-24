package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SendInvitationsCmd = &cobra.Command{
	Use:   "sendinvitations",
	Short: "Send invitation emails",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/user_invitations")
		fmt.Println("OperationID: SendInvitations")
	},
}

func init() {
	Cmd.AddCommand(SendInvitationsCmd)
}
