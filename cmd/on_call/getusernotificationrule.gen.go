package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUserNotificationRuleCmd = &cobra.Command{
	Use:   "getusernotificationrule",
	Short: "Get an On-Call notification rule for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/on-call/users/{user_id}/notification-rules/{rule_id}")
		fmt.Println("OperationID: GetUserNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(GetUserNotificationRuleCmd)
}
