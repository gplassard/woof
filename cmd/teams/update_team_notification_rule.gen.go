package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamNotificationRuleCmd = &cobra.Command{
	Use:     "update-team-notification-rule [team_id] [rule_id]",
	Aliases: []string{"update-notification-rule"},
	Short:   "Update team notification rule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.TeamNotificationRuleRequest{})
		cmdutil.HandleError(err, "failed to update-team-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamNotificationRuleCmd)
}
