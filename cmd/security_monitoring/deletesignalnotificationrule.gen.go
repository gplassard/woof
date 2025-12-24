package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteSignalNotificationRuleCmd = &cobra.Command{
	Use:   "deletesignalnotificationrule",
	Short: "Delete a signal-based notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/security/signals/notification_rules/{id}")
		fmt.Println("OperationID: DeleteSignalNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteSignalNotificationRuleCmd)
}
