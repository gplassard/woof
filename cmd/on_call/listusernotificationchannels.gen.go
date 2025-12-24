package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUserNotificationChannelsCmd = &cobra.Command{
	Use:   "listusernotificationchannels",
	Short: "List On-Call notification channels for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/users/{user_id}/notification-channels")
		fmt.Println("OperationID: ListUserNotificationChannels")
	},
}

func init() {
	Cmd.AddCommand(ListUserNotificationChannelsCmd)
}
