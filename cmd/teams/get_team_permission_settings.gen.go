package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamPermissionSettingsCmd = &cobra.Command{
	Use:     "get-team-permission-settings [team_id]",
	Aliases: []string{"get-permission-settings"},
	Short:   "Get permission settings for a team",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.GetTeamPermissionSettings(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-permission-settings")

		cmd.Println(cmdutil.FormatJSON(res, "team_permission_settings"))
	},
}

func init() {
	Cmd.AddCommand(GetTeamPermissionSettingsCmd)
}
