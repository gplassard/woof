package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTeamPermissionSettingCmd = &cobra.Command{
	Use:     "update-team-permission-setting [team_id] [action]",
	Aliases: []string{"update-permission-setting"},
	Short:   "Update permission setting for team",
	Long: `Update permission setting for team
Documentation: https://docs.datadoghq.com/api/latest/teams/#update-team-permission-setting`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamPermissionSettingResponse
		var err error

		var body datadogV2.TeamPermissionSettingUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateTeamPermissionSetting(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-team-permission-setting")

		fmt.Println(cmdutil.FormatJSON(res, "team_permission_setting"))
	},
}

func init() {

	UpdateTeamPermissionSettingCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTeamPermissionSettingCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTeamPermissionSettingCmd)
}
