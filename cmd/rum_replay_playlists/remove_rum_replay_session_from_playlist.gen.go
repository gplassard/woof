package rum_replay_playlists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var RemoveRumReplaySessionFromPlaylistCmd = &cobra.Command{
	Use: "remove-rum-replay-session-from-playlist [playlist_id] [session_id]",

	Short: "Remove rum replay session from playlist",
	Long: `Remove rum replay session from playlist
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#remove-rum-replay-session-from-playlist`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.RemoveRumReplaySessionFromPlaylist(client.NewContext(apiKey, appKey, site), func() int32 { i, _ := strconv.ParseInt(args[0], 10, 32); return int32(i) }(), args[1])
		cmdutil.HandleError(err, "failed to remove-rum-replay-session-from-playlist")

	},
}

func init() {

	Cmd.AddCommand(RemoveRumReplaySessionFromPlaylistCmd)
}
