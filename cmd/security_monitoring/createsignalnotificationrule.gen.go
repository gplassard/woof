package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateSignalNotificationRuleCmd = &cobra.Command{
	Use:   "createsignalnotificationrule",
	Short: "Create a new signal-based notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/security/signals/notification_rules")
		fmt.Println("OperationID: CreateSignalNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(CreateSignalNotificationRuleCmd)
}
