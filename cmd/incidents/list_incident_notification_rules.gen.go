package incidents

import (
	"fmt"
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
	Long: `List incident notification rules
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-notification-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationRuleArray
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentNotificationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-notification-rules")

		fmt.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {

	Cmd.AddCommand(ListIncidentNotificationRulesCmd)
}
