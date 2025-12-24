package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMonitorNotificationRuleCmd = &cobra.Command{
	Use:   "getmonitornotificationrule",
	Short: "Get a monitor notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/notification_rule/{rule_id}")
		fmt.Println("OperationID: GetMonitorNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorNotificationRuleCmd)
}
