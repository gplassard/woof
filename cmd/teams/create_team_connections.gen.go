package teams

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:     "create-team-connections",
	Aliases: []string{"create-connections"},
	Short:   "Create team connections",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err := api.CreateTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionCreateRequest{})
		cmdutil.HandleError(err, "failed to create-team-connections")

		cmd.Println(cmdutil.FormatJSON(res, "team_connection"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
