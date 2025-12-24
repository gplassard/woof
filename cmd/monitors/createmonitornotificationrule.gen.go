package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "createmonitornotificationrule",
	Short: "Create a monitor notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/monitor/notification_rule")
		fmt.Println("OperationID: CreateMonitorNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(CreateMonitorNotificationRuleCmd)
}
