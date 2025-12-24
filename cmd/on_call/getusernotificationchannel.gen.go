package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserNotificationChannelCmd = &cobra.Command{
	Use:   "getusernotificationchannel",
	Short: "Get an On-Call notification channel for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/users/{user_id}/notification-channels/{channel_id}")
		fmt.Println("OperationID: GetUserNotificationChannel")
	},
}

func init() {
	Cmd.AddCommand(GetUserNotificationChannelCmd)
}
