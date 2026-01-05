package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:     "create-team-connections [payload]",
	Aliases: []string{"create-connections"},
	Short:   "Create team connections",
	Long: `Create team connections
Documentation: https://docs.datadoghq.com/api/latest/teams/#create-team-connections`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamConnectionsResponse
		var err error

		var body datadogV2.TeamConnectionCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		res, _, err = api.CreateTeamConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-team-connections")

		cmd.Println(cmdutil.FormatJSON(res, "team_connection"))
	},
}

func init() {
	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
