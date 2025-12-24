package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteUserNotificationChannelCmd = &cobra.Command{
	Use:   "deleteusernotificationchannel",
	Short: "Delete an On-Call notification channel for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/on-call/users/{user_id}/notification-channels/{channel_id}")
		fmt.Println("OperationID: DeleteUserNotificationChannel")
	},
}

func init() {
	Cmd.AddCommand(DeleteUserNotificationChannelCmd)
}
