package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "createincidentnotificationrule",
	Short: "Create an incident notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/config/notification-rules")
		fmt.Println("OperationID: CreateIncidentNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationRuleCmd)
}
