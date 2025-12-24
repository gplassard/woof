package metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EstimateMetricsOutputSeriesCmd = &cobra.Command{
	Use:   "estimatemetricsoutputseries [metric_name]",
	Short: "Tag Configuration Cardinality Estimator",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.EstimateMetricsOutputSeries(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to estimatemetricsoutputseries: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(EstimateMetricsOutputSeriesCmd)
}
