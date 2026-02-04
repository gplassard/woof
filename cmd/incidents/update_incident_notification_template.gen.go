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

var UpdateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:     "update-incident-notification-template [id]",
	Aliases: []string{"update-notification-template"},
	Short:   "Update incident notification template",
	Long: `Update incident notification template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-notification-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationTemplate
		var err error

		var body datadogV2.PatchIncidentNotificationTemplateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-incident-notification-template")

		fmt.Println(cmdutil.FormatJSON(res, "incident_notification_template"))
	},
}

func init() {

	UpdateIncidentNotificationTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentNotificationTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentNotificationTemplateCmd)
}
