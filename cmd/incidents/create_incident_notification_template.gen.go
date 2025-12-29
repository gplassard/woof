package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "create-incident-notification-template",
	Aliases: []string{ "create-notification-template", },
	Short: "Create incident notification template",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), datadogV2.CreateIncidentNotificationTemplateRequest{})
		cmdutil.HandleError(err, "failed to create-incident-notification-template")

		cmd.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
