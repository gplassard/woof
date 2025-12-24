package monitors

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMonitorNotificationRulesCmd = &cobra.Command{
	Use:   "getmonitornotificationrules",
	Short: "Get all monitor notification rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/monitor/notification_rule")
		fmt.Println("OperationID: GetMonitorNotificationRules")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorNotificationRulesCmd)
}
