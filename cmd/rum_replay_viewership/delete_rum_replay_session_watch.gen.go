package rum_replay_viewership

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRumReplaySessionWatchCmd = &cobra.Command{
	Use: "delete-rum-replay-session-watch [session_id]",

	Short: "Delete rum replay session watch",
	Long: `Delete rum replay session watch
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-viewership/#delete-rum-replay-session-watch`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRumReplayViewershipApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteRumReplaySessionWatch(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-rum-replay-session-watch")

	},
}

func init() {

	Cmd.AddCommand(DeleteRumReplaySessionWatchCmd)
}
