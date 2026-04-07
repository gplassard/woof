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

var ListRumReplayPlaylistSessionsCmd = &cobra.Command{
	Use:     "list-rum-replay-playlist-sessions [playlist_id]",
	Aliases: []string{"list-sessions"},
	Short:   "List rum replay playlist sessions",
	Long: `List rum replay playlist sessions
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-playlists/#list-rum-replay-playlist-sessions`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PlaylistsSessionArray
		var err error

		api := datadogV2.NewRumReplayPlaylistsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRumReplayPlaylistSessions(client.NewContext(apiKey, appKey, site), func() int32 { i, _ := strconv.ParseInt(args[0], 10, 32); return int32(i) }())
		cmdutil.HandleError(err, "failed to list-rum-replay-playlist-sessions")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_session"))
	},
}

func init() {

	Cmd.AddCommand(ListRumReplayPlaylistSessionsCmd)
}
