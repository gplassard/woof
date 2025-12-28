package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "get-incident-notification-template [id]",
	Short: "Get incident notification template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to get-incident-notification-template: %v", err)
		}

		cmdutil.PrintJSON(res, "notification_templates")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentNotificationTemplateCmd)
}
