package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "create-incident-notification-rule",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create an incident notification rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.CreateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateIncidentNotificationRuleRequest{})
		cmdutil.HandleError(err, "failed to create-incident-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationRuleCmd)
}
