package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTeamConnectionsCmd = &cobra.Command{
	Use:     "list-team-connections",
	Aliases: []string{"list-connections"},
	Short:   "List team connections",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamConnectionsResponse
		var err error

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.ListTeamConnections(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-team-connections")

		cmd.Println(cmdutil.FormatJSON(res, "team_connection"))
	},
}

func init() {
	Cmd.AddCommand(ListTeamConnectionsCmd)
}
