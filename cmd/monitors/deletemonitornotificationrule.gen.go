package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "deletemonitornotificationrule",
	Short: "Delete a monitor notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/monitor/notification_rule/{rule_id}")
		fmt.Println("OperationID: DeleteMonitorNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorNotificationRuleCmd)
}
