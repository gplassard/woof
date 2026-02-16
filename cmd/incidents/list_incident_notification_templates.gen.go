package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIncidentNotificationTemplatesCmd = &cobra.Command{
	Use:     "list-incident-notification-templates",
	Aliases: []string{"list-notification-templates"},
	Short:   "List incident notification templates",
	Long: `List incident notification templates
Documentation: https://docs.datadoghq.com/api/latest/incidents/#list-incident-notification-templates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationTemplateArray
		var err error

		optionalParams := datadogV2.NewListIncidentNotificationTemplatesOptionalParameters()

		if cmd.Flags().Changed("filter-incident-type") {
			val, _ := cmd.Flags().GetString("filter-incident-type")
			optionalParams.WithFilterIncidentType(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentNotificationTemplates(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-incident-notification-templates")

		fmt.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {

	ListIncidentNotificationTemplatesCmd.Flags().String("filter-incident-type", "", "Optional incident type ID filter.")

	ListIncidentNotificationTemplatesCmd.Flags().String("include", "", "Comma-separated list of relationships to include. Supported values: 'created_by_user', 'last_modified_by_user', 'incident_type'")

	Cmd.AddCommand(ListIncidentNotificationTemplatesCmd)
}
