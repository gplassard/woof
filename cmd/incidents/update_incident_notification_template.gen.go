package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var UpdateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:   "update-incident-notification-template [id]",
	Aliases: []string{ "update-notification-template", },
	Short: "Update incident notification template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.PatchIncidentNotificationTemplateRequest{})
		cmdutil.HandleError(err, "failed to update-incident-notification-template")

		cmdutil.PrintJSON(res, "notification_templates")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationTemplateCmd)
}
