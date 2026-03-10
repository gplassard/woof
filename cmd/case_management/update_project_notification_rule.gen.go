package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateProjectNotificationRuleCmd = &cobra.Command{
	Use: "update-project-notification-rule [project_id] [notification_rule_id]",

	Short: "Update a notification rule",
	Long: `Update a notification rule
Documentation: https://docs.datadoghq.com/api/latest/case-management/#update-project-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.CaseNotificationRuleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UpdateProjectNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-project-notification-rule")

	},
}

func init() {

	UpdateProjectNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateProjectNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateProjectNotificationRuleCmd)
}
