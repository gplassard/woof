package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentNotificationTemplatesCmd = &cobra.Command{
	Use:   "listincidentnotificationtemplates",
	Short: "List incident notification templates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/config/notification-templates")
		fmt.Println("OperationID: ListIncidentNotificationTemplates")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentNotificationTemplatesCmd)
}
