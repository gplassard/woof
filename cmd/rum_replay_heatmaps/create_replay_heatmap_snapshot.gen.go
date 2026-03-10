package rum_replay_heatmaps

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateReplayHeatmapSnapshotCmd = &cobra.Command{
	Use: "create-replay-heatmap-snapshot",

	Short: "Create replay heatmap snapshot",
	Long: `Create replay heatmap snapshot
Documentation: https://docs.datadoghq.com/api/latest/rum-replay-heatmaps/#create-replay-heatmap-snapshot`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Snapshot
		var err error

		var body datadogV2.SnapshotCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumReplayHeatmapsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateReplayHeatmapSnapshot(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-replay-heatmap-snapshot")

		fmt.Println(cmdutil.FormatJSON(res, "snapshots"))
	},
}

func init() {

	CreateReplayHeatmapSnapshotCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateReplayHeatmapSnapshotCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateReplayHeatmapSnapshotCmd)
}
