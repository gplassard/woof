package rum_replay_playlists

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var BulkRemoveRumReplayPlaylistSessionsCmd = &cobra.Command{
	Use:     "bulk-remove-rum-replay-playlist-sessions [playlist_id]",
	Aliases: []string{"bulk-remove-sessions"},
	Short:   "Bulk remove rum replay playlist sessions",
	Long: `Bulk remove rum replay playlist sessions
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#bulk-remove-rum-replay-playlist-sessions`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.SessionIdArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.BulkRemoveRumReplayPlaylistSessions(client.NewContext(apiKey, appKey, site), func() int32 { i, _ := strconv.ParseInt(args[0], 10, 32); return int32(i) }(), body)
		cmdutil.HandleError(err, "failed to bulk-remove-rum-replay-playlist-sessions")

	},
}

func init() {

	BulkRemoveRumReplayPlaylistSessionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	BulkRemoveRumReplayPlaylistSessionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(BulkRemoveRumReplayPlaylistSessionsCmd)
}
