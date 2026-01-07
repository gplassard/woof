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
	Long: `List incident notification rules
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-notification-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationRuleArray
		var err error

		optionalParams := datadogV2.NewListIncidentNotificationRulesOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentNotificationRules(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {

	ListIncidentNotificationRulesCmd.Flags().String("include", "", "Comma-separated list of resources to include. Supported values: 'created_by_user', 'last_modified_by_user', 'incident_type', 'notification_template'")

	Cmd.AddCommand(ListIncidentNotificationRulesCmd)
}
