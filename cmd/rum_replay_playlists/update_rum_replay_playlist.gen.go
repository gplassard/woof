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

var UpdateRumReplayPlaylistCmd = &cobra.Command{
	Use:     "update-rum-replay-playlist [playlist_id]",
	Aliases: []string{"update"},
	Short:   "Update rum replay playlist",
	Long: `Update rum replay playlist
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#update-rum-replay-playlist`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Playlist
		var err error

		var body datadogV2.Playlist
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRumReplayPlaylist(client.NewContext(apiKey, appKey, site), func() int32 { i, _ := strconv.ParseInt(args[0], 10, 32); return int32(i) }(), body)
		cmdutil.HandleError(err, "failed to update-rum-replay-playlist")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_playlist"))
	},
}

func init() {

	UpdateRumReplayPlaylistCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRumReplayPlaylistCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRumReplayPlaylistCmd)
}
