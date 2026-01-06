package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTeamConnectionsCmd = &cobra.Command{
	Use:     "create-team-connections",
	Aliases: []string{"create-connections"},
	Short:   "Create team connections",
	Long: `Create team connections
Documentation: https://docs.datadoghq.com/api/latest/teams/#create-team-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TeamConnectionsResponse
		var err error

		var body datadogV2.TeamConnectionCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateTeamConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-team-connections")

		cmd.Println(cmdutil.FormatJSON(res, "team_connection"))
	},
}

func init() {

	CreateTeamConnectionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTeamConnectionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTeamConnectionsCmd)
}
