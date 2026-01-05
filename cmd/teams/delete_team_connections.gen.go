package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTeamConnectionsCmd = &cobra.Command{
	Use:     "delete-team-connections",
	Aliases: []string{"delete-connections"},
	Short:   "Delete team connections",
	Long: `Delete team connections
Documentation: https://docs.datadoghq.com/api/latest/teams/#delete-team-connections`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.TeamConnectionDeleteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err = api.DeleteTeamConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-team-connections")

	},
}

func init() {

	DeleteTeamConnectionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteTeamConnectionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
