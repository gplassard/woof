package rum_replay_heatmaps

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteReplayHeatmapSnapshotCmd = &cobra.Command{
	Use: "delete-replay-heatmap-snapshot [snapshot_id]",

	Short: "Delete replay heatmap snapshot",
	Long: `Delete replay heatmap snapshot
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-heatmaps/#delete-replay-heatmap-snapshot`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRumReplayHeatmapsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteReplayHeatmapSnapshot(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-replay-heatmap-snapshot")

	},
}

func init() {

	Cmd.AddCommand(DeleteReplayHeatmapSnapshotCmd)
}
