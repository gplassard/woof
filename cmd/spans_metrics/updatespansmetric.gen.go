package spans_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateSpansMetricCmd = &cobra.Command{
	Use:   "updatespansmetric [metric_id]",
	Short: "Update a span-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateSpansMetric(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SpansMetricUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatespansmetric: %v", err)
		}

		cmdutil.PrintJSON(res, "spans_metrics")
	},
}

func init() {
	Cmd.AddCommand(UpdateSpansMetricCmd)
}
