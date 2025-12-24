package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "createincidentnotificationtemplate",
	Short: "Create incident notification template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/config/notification-templates")
		fmt.Println("OperationID: CreateIncidentNotificationTemplate")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
