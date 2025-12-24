package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "getincidentnotificationtemplate",
	Short: "Get incident notification template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/notification-templates/{id}")
		fmt.Println("OperationID: GetIncidentNotificationTemplate")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentNotificationTemplateCmd)
}
