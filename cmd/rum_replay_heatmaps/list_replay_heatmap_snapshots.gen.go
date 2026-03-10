package rum_replay_heatmaps

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListReplayHeatmapSnapshotsCmd = &cobra.Command{
	Use: "list-replay-heatmap-snapshots [filter[view_name]]",

	Short: "List replay heatmap snapshots",
	Long: `List replay heatmap snapshots
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-heatmaps/#list-replay-heatmap-snapshots`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SnapshotArray
		var err error

		api := datadogV2.NewRumReplayHeatmapsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListReplayHeatmapSnapshots(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-replay-heatmap-snapshots")

		fmt.Println(cmdutil.FormatJSON(res, "snapshots"))
	},
}

func init() {

	Cmd.AddCommand(ListReplayHeatmapSnapshotsCmd)
}
