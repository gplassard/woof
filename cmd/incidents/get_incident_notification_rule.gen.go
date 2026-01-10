package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "get-incident-notification-rule [id]",
	Aliases: []string{"get-notification-rule"},
	Short:   "Get an incident notification rule",
	Long: `Get an incident notification rule
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-incident-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationRule
		var err error

		optionalParams := datadogV2.NewGetIncidentNotificationRuleOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), *optionalParams)
		cmdutil.HandleError(err, "failed to get-incident-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "incident_notification_rules"))
	},
}

func init() {

	GetIncidentNotificationRuleCmd.Flags().String("include", "", "Comma-separated list of resources to include. Supported values: 'created_by_user', 'last_modified_by_user', 'incident_type', 'notification_template'")

	Cmd.AddCommand(GetIncidentNotificationRuleCmd)
}
