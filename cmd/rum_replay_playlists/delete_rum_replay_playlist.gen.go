package rum_replay_playlists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var DeleteRumReplayPlaylistCmd = &cobra.Command{
	Use:     "delete-rum-replay-playlist [playlist_id]",
	Aliases: []string{"delete"},
	Short:   "Delete rum replay playlist",
	Long: `Delete rum replay playlist
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#delete-rum-replay-playlist`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteRumReplayPlaylist(client.NewContext(apiKey, appKey, site), func() int32 { i, _ := strconv.ParseInt(args[0], 10, 32); return int32(i) }())
		cmdutil.HandleError(err, "failed to delete-rum-replay-playlist")

	},
}

func init() {

	Cmd.AddCommand(DeleteRumReplayPlaylistCmd)
}
