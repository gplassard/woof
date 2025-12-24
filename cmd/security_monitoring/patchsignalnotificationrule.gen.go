package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var PatchSignalNotificationRuleCmd = &cobra.Command{
	Use:   "patchsignalnotificationrule",
	Short: "Patch a signal-based notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/security/signals/notification_rules/{id}")
		fmt.Println("OperationID: PatchSignalNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(PatchSignalNotificationRuleCmd)
}
