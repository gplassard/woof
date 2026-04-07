package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamNotificationRuleCmd = &cobra.Command{
	Use:     "get-team-notification-rule [team_id] [rule_id]",
	Aliases: []string{"get-notification-rule"},
	Short:   "Get team notification rule",
	Long: `Get team notification rule
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamNotificationRuleResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-team-notification-rule")

		fmt.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamNotificationRuleCmd)
}
