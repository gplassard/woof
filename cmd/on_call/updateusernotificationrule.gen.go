package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateUserNotificationRuleCmd = &cobra.Command{
	Use:   "updateusernotificationrule",
	Short: "Update an On-Call notification rule for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/on-call/users/{user_id}/notification-rules/{rule_id}")
		fmt.Println("OperationID: UpdateUserNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateUserNotificationRuleCmd)
}
