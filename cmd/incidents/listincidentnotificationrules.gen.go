package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentNotificationRulesCmd = &cobra.Command{
	Use:   "listincidentnotificationrules",
	Short: "List incident notification rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/notification-rules")
		fmt.Println("OperationID: ListIncidentNotificationRules")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentNotificationRulesCmd)
}
