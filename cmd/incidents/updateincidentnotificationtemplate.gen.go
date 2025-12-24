package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "updateincidentnotificationtemplate",
	Short: "Update incident notification template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/config/notification-templates/{id}")
		fmt.Println("OperationID: UpdateIncidentNotificationTemplate")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationTemplateCmd)
}
