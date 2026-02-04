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

var UpdateIncidentNotificationRuleCmd = &cobra.Command{
	Use:     "update-incident-notification-rule [id]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update an incident notification rule",
	Long: `Update an incident notification rule
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-incident-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentNotificationRule
		var err error

		var body datadogV2.PutIncidentNotificationRuleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateIncidentNotificationRule(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-incident-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "incident_notification_rule"))
	},
}

func init() {

	UpdateIncidentNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateIncidentNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateIncidentNotificationRuleCmd)
}
