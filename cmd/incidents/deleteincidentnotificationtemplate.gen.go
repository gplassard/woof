package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "deleteincidentnotificationtemplate",
	Short: "Delete a notification template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/config/notification-templates/{id}")
		fmt.Println("OperationID: DeleteIncidentNotificationTemplate")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentNotificationTemplateCmd)
}
