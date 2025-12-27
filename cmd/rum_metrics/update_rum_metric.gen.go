package rum_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateRumMetricCmd = &cobra.Command{
	Use:   "update_rum_metric [metric_id]",
	Short: "Update a rum-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateRumMetric(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RumMetricUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_rum_metric: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_metrics")
	},
}

func init() {
	Cmd.AddCommand(UpdateRumMetricCmd)
}
