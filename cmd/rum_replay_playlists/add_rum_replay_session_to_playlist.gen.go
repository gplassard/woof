package rum_replay_playlists

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var AddRumReplaySessionToPlaylistCmd = &cobra.Command{
	Use: "add-rum-replay-session-to-playlist [ts] [playlist_id] [session_id]",

	Short: "Add rum replay session to playlist",
	Long: `Add rum replay session to playlist
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#add-rum-replay-session-to-playlist`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PlaylistsSession
		var err error

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AddRumReplaySessionToPlaylist(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), func() int32 { i, _ := strconv.ParseInt(args[1], 10, 32); return int32(i) }(), args[2])
		cmdutil.HandleError(err, "failed to add-rum-replay-session-to-playlist")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_session"))
	},
}

func init() {

	Cmd.AddCommand(AddRumReplaySessionToPlaylistCmd)
}
