package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:     "update-incident-notification-template [id] [payload]",
	Aliases: []string{"update-notification-template"},
	Short:   "Update incident notification template",
	Long: `Update incident notification template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-notification-template`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationTemplate
		var err error

		var body datadogV2.PatchIncidentNotificationTemplateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.UpdateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-incident-notification-template")

		cmd.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationTemplateCmd)
}
