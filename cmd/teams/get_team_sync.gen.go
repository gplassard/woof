package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTeamSyncCmd = &cobra.Command{
	Use:     "get-team-sync [filter[source]]",
	Aliases: []string{"get-sync"},
	Short:   "Get team sync configurations",
	Long: `Get team sync configurations
Documentation: https://docs.datadoghq.com/api/latest/teams/#get-team-sync`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamSyncResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTeamSync(client.NewContext(apiKey, appKey, site), datadogV2.TeamSyncAttributesSource(args[0]))
		cmdutil.HandleError(err, "failed to get-team-sync")

		cmd.Println(cmdutil.FormatJSON(res, "team_sync_bulk"))
	},
}

func init() {

	Cmd.AddCommand(GetTeamSyncCmd)
}
