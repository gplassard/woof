package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var DeleteTeamConnectionsCmd = &cobra.Command{
	Use:     "delete-team-connections [payload]",
	Aliases: []string{"delete-connections"},
	Short:   "Delete team connections",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.TeamConnectionDeleteRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamConnections(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-team-connections")

	},
}

func init() {
	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
