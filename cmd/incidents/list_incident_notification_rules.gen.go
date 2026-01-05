package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentNotificationRulesCmd = &cobra.Command{
	Use:     "list-incident-notification-rules",
	Aliases: []string{"list-notification-rules"},
	Short:   "List incident notification rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.ListIncidentNotificationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(ListIncidentNotificationRulesCmd)
}
