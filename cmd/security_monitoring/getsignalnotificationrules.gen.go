package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSignalNotificationRulesCmd = &cobra.Command{
	Use:   "getsignalnotificationrules",
	Short: "Get the list of signal-based notification rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/signals/notification_rules")
		fmt.Println("OperationID: GetSignalNotificationRules")
	},
}

func init() {
	Cmd.AddCommand(GetSignalNotificationRulesCmd)
}
