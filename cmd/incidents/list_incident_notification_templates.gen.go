package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListIncidentNotificationTemplatesCmd = &cobra.Command{
	Use:   "list_incident_notification_templates",
	Short: "List incident notification templates",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentNotificationTemplates(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_incident_notification_templates: %v", err)
		}

		cmdutil.PrintJSON(res, "notification_templates")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentNotificationTemplatesCmd)
}
