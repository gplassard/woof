package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateUserNotificationChannelCmd = &cobra.Command{
	Use:   "createusernotificationchannel",
	Short: "Create an On-Call notification channel for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/users/{user_id}/notification-channels")
		fmt.Println("OperationID: CreateUserNotificationChannel")
	},
}

func init() {
	Cmd.AddCommand(CreateUserNotificationChannelCmd)
}
