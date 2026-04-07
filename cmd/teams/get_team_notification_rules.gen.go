package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamNotificationRulesCmd = &cobra.Command{
	Use:     "get-team-notification-rules [team_id]",
	Aliases: []string{"get-notification-rules"},
	Short:   "Get team notification rules",
	Long: `Get team notification rules
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-notification-rules`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamNotificationRulesResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamNotificationRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-notification-rules")

		fmt.Println(cmdutil.FormatJSON(res, "team_notification_rules"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamNotificationRulesCmd)
}
