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

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIncidentNotificationTemplates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-incident-notification-templates")

		fmt.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {

	Cmd.AddCommand(ListIncidentNotificationTemplatesCmd)
}
