package rum_replay_viewership

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRumReplaySessionWatchCmd = &cobra.Command{
	Use: "create-rum-replay-session-watch [session_id]",

	Short: "Create rum replay session watch",
	Long: `Create rum replay session watch
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-viewership/#create-rum-replay-session-watch`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Watch
		var err error

		var body datadogV2.Watch
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayViewershipApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRumReplaySessionWatch(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-rum-replay-session-watch")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_watch"))
	},
}

func init() {

	CreateRumReplaySessionWatchCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRumReplaySessionWatchCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRumReplaySessionWatchCmd)
}
