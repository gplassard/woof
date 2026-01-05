package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateTeamPermissionSettingCmd = &cobra.Command{
	Use:     "update-team-permission-setting [team_id] [action] [payload]",
	Aliases: []string{"update-permission-setting"},
	Short:   "Update permission setting for team",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamPermissionSettingResponse
		var err error

		var body datadogV2.TeamPermissionSettingUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.UpdateTeamPermissionSetting(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-permission-setting")

		cmd.Println(cmdutil.FormatJSON(res, "team_permission_settings"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTeamPermissionSettingCmd)
}
