package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "getincidentnotificationrule",
	Short: "Get an incident notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/notification-rules/{id}")
		fmt.Println("OperationID: GetIncidentNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentNotificationRuleCmd)
}
