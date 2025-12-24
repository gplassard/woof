package on_call

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateUserNotificationRuleCmd = &cobra.Command{
	Use:   "createusernotificationrule",
	Short: "Create an On-Call notification rule for a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/users/{user_id}/notification-rules")
		fmt.Println("OperationID: CreateUserNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(CreateUserNotificationRuleCmd)
}
