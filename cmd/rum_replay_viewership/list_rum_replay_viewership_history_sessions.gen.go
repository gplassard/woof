package rum_replay_viewership

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRumReplayViewershipHistorySessionsCmd = &cobra.Command{
	Use:     "list-rum-replay-viewership-history-sessions",
	Aliases: []string{"list-history-sessions"},
	Short:   "List rum replay viewership history sessions",
	Long: `List rum replay viewership history sessions
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-viewership/#list-rum-replay-viewership-history-sessions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ViewershipHistorySessionArray
		var err error

		api := datadogV2.NewRumReplayViewershipApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRumReplayViewershipHistorySessions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-rum-replay-viewership-history-sessions")

		fmt.Println(cmdutil.FormatJSON(res, "rum_replay_session"))
	},
}

func init() {

	Cmd.AddCommand(ListRumReplayViewershipHistorySessionsCmd)
}
