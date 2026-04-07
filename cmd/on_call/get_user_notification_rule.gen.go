package on_call

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUserNotificationRuleCmd = &cobra.Command{
	Use: "get-user-notification-rule [user_id] [rule_id]",

	Short: "Get an On-Call notification rule for a user",
	Long: `Get an On-Call notification rule for a user
Documentation: https://docs.datadoghq.com/api/latest/on-call/#get-user-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OnCallNotificationRule
		var err error

		api := datadogV2.NewOnCallApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetUserNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-user-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "notification_rules"))
	},
}

func init() {

	Cmd.AddCommand(GetUserNotificationRuleCmd)
}
