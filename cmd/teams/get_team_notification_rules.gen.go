package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamNotificationRulesCmd = &cobra.Command{
	Use:     "get-team-notification-rules [team_id]",
	Aliases: []string{"get-notification-rules"},
	Short:   "Get team notification rules",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamNotificationRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-notification-rules")

		cmd.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {
	Cmd.AddCommand(GetTeamNotificationRulesCmd)
}
