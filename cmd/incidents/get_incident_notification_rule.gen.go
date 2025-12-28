package incidents

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var GetIncidentNotificationRuleCmd = &cobra.Command{
	Use:   "get-incident-notification-rule [id]",
	Aliases: []string{ "get-notification-rule", },
	Short: "Get an incident notification rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.GetIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-incident-notification-rule")

		cmdutil.PrintJSON(res, "incident_notification_rules")
	},
}

func init() {
	Cmd.AddCommand(GetIncidentNotificationRuleCmd)
}
