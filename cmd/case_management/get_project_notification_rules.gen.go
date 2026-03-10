package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectNotificationRulesCmd = &cobra.Command{
	Use: "get-project-notification-rules [project_id]",

	Short: "Get notification rules",
	Long: `Get notification rules
Documentation: https://docs.datadoghq.com/api/latest/case-management/#get-project-notification-rules`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseNotificationRulesResponse
		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetProjectNotificationRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-project-notification-rules")

		fmt.Println(cmdutil.FormatJSON(res, "notification_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetProjectNotificationRulesCmd)
}
