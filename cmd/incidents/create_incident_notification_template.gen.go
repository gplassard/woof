package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "create_incident_notification_template",
	Short: "Create incident notification template",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), datadogV2.CreateIncidentNotificationTemplateRequest{})
		if err != nil {
			log.Fatalf("failed to create_incident_notification_template: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
