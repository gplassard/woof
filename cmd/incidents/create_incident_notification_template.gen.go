package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateIncidentNotificationTemplateCmd = &cobra.Command{
	Use:     "create-incident-notification-template",
	Aliases: []string{"create-notification-template"},
	Short:   "Create incident notification template",
	Long: `Create incident notification template
Documentation: https://docs.datadoghq.com/api/latest/incidents/#create-incident-notification-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationTemplate
		var err error

		var body datadogV2.CreateIncidentNotificationTemplateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateIncidentNotificationTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-incident-notification-template")

		fmt.Println(cmdutil.FormatJSON(res, "notification_templates"))
	},
}

func init() {

	CreateIncidentNotificationTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateIncidentNotificationTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateIncidentNotificationTemplateCmd)
}
