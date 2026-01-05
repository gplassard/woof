package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SyncTeamsCmd = &cobra.Command{
	Use:     "sync-teams [payload]",
	Aliases: []string{"sync"},
	Short:   "Link Teams with GitHub Teams",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.TeamSyncRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.SyncTeams(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to sync-teams")

	},
}

func init() {
	Cmd.AddCommand(SyncTeamsCmd)
}
