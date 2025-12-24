package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSignalNotificationRuleCmd = &cobra.Command{
	Use:   "getsignalnotificationrule",
	Short: "Get details of a signal-based notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/signals/notification_rules/{id}")
		fmt.Println("OperationID: GetSignalNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(GetSignalNotificationRuleCmd)
}
