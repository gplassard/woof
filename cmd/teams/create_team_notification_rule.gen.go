package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamNotificationRuleCmd = &cobra.Command{
	Use:     "create-team-notification-rule [team_id]",
	Aliases: []string{"create-notification-rule"},
	Short:   "Create team notification rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.TeamNotificationRuleRequest{})
		cmdutil.HandleError(err, "failed to create-team-notification-rule")

		cmd.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamNotificationRuleCmd)
}
