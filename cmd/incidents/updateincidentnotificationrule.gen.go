package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "updateincidentnotificationrule",
	Short: "Update an incident notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/incidents/config/notification-rules/{id}")
		fmt.Println("OperationID: UpdateIncidentNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationRuleCmd)
}
