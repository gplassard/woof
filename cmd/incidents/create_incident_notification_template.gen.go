package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:     "create-incident-notification-template [payload]",
	Aliases: []string{"create-notification-template"},
	Short:   "Create incident notification template",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CreateIncidentNotificationTemplateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-notification-template")

		cmd.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
