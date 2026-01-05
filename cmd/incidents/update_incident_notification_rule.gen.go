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

var UpdateIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "update-incident-notification-rule [id] [payload]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update an incident notification rule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.PutIncidentNotificationRuleRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-incident-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationRuleCmd)
}
