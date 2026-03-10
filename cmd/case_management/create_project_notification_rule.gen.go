package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateProjectNotificationRuleCmd = &cobra.Command{
	Use: "create-project-notification-rule [project_id]",

	Short: "Create a notification rule",
	Long: `Create a notification rule
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-project-notification-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseNotificationRuleResponse
		var err error

		var body datadogV2.CaseNotificationRuleCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateProjectNotificationRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-project-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "notification_rule"))
	},
}

func init() {

	CreateProjectNotificationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateProjectNotificationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateProjectNotificationRuleCmd)
}
