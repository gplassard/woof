package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "delete-incident-notification-rule [id]",
	Aliases: []string{"delete-notification-rule"},
	Short:   "Delete an incident notification rule",
	Long: `Delete an incident notification rule
Documentation: https://docs.datadoghq.com/api/latest/incidents/#delete-incident-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		optionalParams := datadogV2.NewDeleteIncidentNotificationRuleOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), *optionalParams)
		cmdutil.HandleError(err, "failed to delete-incident-notification-rule")

	},
}

func init() {

	DeleteIncidentNotificationRuleCmd.Flags().String("include", "", "Comma-separated list of resources to include. Supported values: 'created_by_user', 'last_modified_by_user', 'incident_type', 'notification_template'")

	Cmd.AddCommand(DeleteIncidentNotificationRuleCmd)
}
