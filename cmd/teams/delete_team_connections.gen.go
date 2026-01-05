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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.DeleteTeamConnections(client.NewContext(apiKey, appKey, site), datadogV2.TeamConnectionDeleteRequest{})
		cmdutil.HandleError(err, "failed to delete-team-connections")

	},
}

func init() {
	Cmd.AddCommand(DeleteTeamConnectionsCmd)
}
