package rum_replay_sessions

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSegmentsCmd = &cobra.Command{
	Use: "get-segments [view_id] [session_id]",

	Short: "Get segments",
	Long: `Get segments
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-sessions/#get-segments`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRumReplaySessionsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.GetSegments(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-segments")

	},
}

func init() {

	Cmd.AddCommand(GetSegmentsCmd)
}
