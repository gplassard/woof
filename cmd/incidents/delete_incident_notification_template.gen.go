package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteIncidentNotificationTemplateCmd = &cobra.Command{
	Use:     "delete-incident-notification-template [id]",
	Aliases: []string{"delete-notification-template"},
	Short:   "Delete a notification template",
	Long: `Delete a notification template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#delete-incident-notification-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-incident-notification-template")

	},
}

func init() {

	Cmd.AddCommand(DeleteIncidentNotificationTemplateCmd)
}
