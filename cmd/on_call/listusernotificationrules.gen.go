package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUserNotificationRulesCmd = &cobra.Command{
	Use:   "listusernotificationrules",
	Short: "List On-Call notification rules for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/users/{user_id}/notification-rules")
		fmt.Println("OperationID: ListUserNotificationRules")
	},
}

func init() {
	Cmd.AddCommand(ListUserNotificationRulesCmd)
}
