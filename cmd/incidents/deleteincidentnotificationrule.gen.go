package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "deleteincidentnotificationrule",
	Short: "Delete an incident notification rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/config/notification-rules/{id}")
		fmt.Println("OperationID: DeleteIncidentNotificationRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentNotificationRuleCmd)
}
