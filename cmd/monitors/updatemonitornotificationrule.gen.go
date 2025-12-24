package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "updatemonitornotificationrule",
	Short: "Update a monitor notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/monitor/notification_rule/{rule_id}")
		fmt.Println("OperationID: UpdateMonitorNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateMonitorNotificationRuleCmd)
}
