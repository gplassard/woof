package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamPermissionSettingCmd = &cobra.Command{
	Use:     "update-team-permission-setting [team_id] [action]",
	Aliases: []string{"update-permission-setting"},
	Short:   "Update permission setting for team",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.UpdateTeamPermissionSetting(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.TeamPermissionSettingUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-team-permission-setting")

		cmd.Println(cmdutil.FormatJSON(res, "team_permission_settings"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamPermissionSettingCmd)
}
