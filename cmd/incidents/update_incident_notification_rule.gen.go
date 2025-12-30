package incidents

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "update-incident-notification-rule [id]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update an incident notification rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.UpdateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), datadogV2.PutIncidentNotificationRuleRequest{})
		cmdutil.HandleError(err, "failed to update-incident-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentNotificationRuleCmd)
}
