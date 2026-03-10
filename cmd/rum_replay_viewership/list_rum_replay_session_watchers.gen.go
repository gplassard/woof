package rum_replay_viewership

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRumReplaySessionWatchersCmd = &cobra.Command{
	Use: "list-rum-replay-session-watchers [session_id]",

	Short: "List rum replay session watchers",
	Long: `List rum replay session watchers
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-viewership/#list-rum-replay-session-watchers`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WatcherArray
		var err error

		api := datadogV2.NewRumReplayViewershipApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRumReplaySessionWatchers(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-rum-replay-session-watchers")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_watcher"))
	},
}

func init() {

	Cmd.AddCommand(ListRumReplaySessionWatchersCmd)
}
