package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteUserNotificationRuleCmd = &cobra.Command{
	Use:   "deleteusernotificationrule",
	Short: "Delete an On-Call notification rule for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/on-call/users/{user_id}/notification-rules/{rule_id}")
		fmt.Println("OperationID: DeleteUserNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteUserNotificationRuleCmd)
}
