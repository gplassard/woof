package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var DeleteIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "delete-incident-notification-template [id]",
	Short: "Delete a notification template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		_, err := api.DeleteIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to delete-incident-notification-template: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentNotificationTemplateCmd)
}
