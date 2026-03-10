package rum_replay_heatmaps

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateReplayHeatmapSnapshotCmd = &cobra.Command{
	Use: "update-replay-heatmap-snapshot [snapshot_id]",

	Short: "Update replay heatmap snapshot",
	Long: `Update replay heatmap snapshot
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-heatmaps/#update-replay-heatmap-snapshot`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Snapshot
		var err error

		var body datadogV2.SnapshotUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayHeatmapsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateReplayHeatmapSnapshot(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-replay-heatmap-snapshot")

		fmt.Println(cmdutil.FormatJSON(res, "snapshots"))
	},
}

func init() {

	UpdateReplayHeatmapSnapshotCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateReplayHeatmapSnapshotCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateReplayHeatmapSnapshotCmd)
}
