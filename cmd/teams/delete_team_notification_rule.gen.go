package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamNotificationRuleCmd = &cobra.Command{
	Use:     "delete-team-notification-rule [team_id] [rule_id]",
	Aliases: []string{"delete-notification-rule"},
	Short:   "Delete team notification rule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-team-notification-rule")

	},
}

func init() {
	Cmd.AddCommand(DeleteTeamNotificationRuleCmd)
}
