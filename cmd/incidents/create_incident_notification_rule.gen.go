package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "create-incident-notification-rule [payload]",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create an incident notification rule",
	Long: `Create an incident notification rule
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationRule
		var err error

		var body datadogV2.CreateIncidentNotificationRuleRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err = api.CreateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentNotificationRuleCmd)
}
