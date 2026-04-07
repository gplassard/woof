package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamNotificationRuleCmd = &cobra.Command{
	Use:     "delete-team-notification-rule [team_id] [rule_id]",
	Aliases: []string{"delete-notification-rule"},
	Short:   "Delete team notification rule",
	Long: `Delete team notification rule
Documentation: https://docs.datadoghq.com/api/latest/teams/#delete-team-notification-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteTeamNotificationRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-team-notification-rule")

	},
}

func init() {

	Cmd.AddCommand(DeleteTeamNotificationRuleCmd)
}
