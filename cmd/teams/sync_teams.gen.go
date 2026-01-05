package teams

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SyncTeamsCmd = &cobra.Command{
	Use:     "sync-teams",
	Aliases: []string{"sync"},
	Short:   "Link Teams with GitHub Teams",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTeamsApi(client.NewAPIClient())
		_, err := api.SyncTeams(client.NewContext(apiKey, appKey, site), datadogV2.TeamSyncRequest{})
		cmdutil.HandleError(err, "failed to sync-teams")

	},
}

func init() {
	Cmd.AddCommand(SyncTeamsCmd)
}
