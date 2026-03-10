package rum_replay_playlists

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRumReplayPlaylistCmd = &cobra.Command{
	Use:     "create-rum-replay-playlist",
	Aliases: []string{"create"},
	Short:   "Create rum replay playlist",
	Long: `Create rum replay playlist
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#create-rum-replay-playlist`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Playlist
		var err error

		var body datadogV2.Playlist
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRumReplayPlaylist(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-rum-replay-playlist")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_playlist"))
	},
}

func init() {

	CreateRumReplayPlaylistCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRumReplayPlaylistCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRumReplayPlaylistCmd)
}
