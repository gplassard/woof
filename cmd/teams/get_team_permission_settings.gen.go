package teams

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamPermissionSettingsCmd = &cobra.Command{
	Use:     "get-team-permission-settings [team_id]",
	Aliases: []string{"get-permission-settings"},
	Short:   "Get permission settings for a team",
	Long: `Get permission settings for a team
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-permission-settings`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamPermissionSettingsResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamPermissionSettings(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-team-permission-settings")

		fmt.Println(cmdutil.FormatJSON(res, "team_permission_setting"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamPermissionSettingsCmd)
}
