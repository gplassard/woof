package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EstimateMetricsOutputSeriesCmd = &cobra.Command{
	Use:   "estimate-metrics-output-series [metric_name]",
	Aliases: []string{ "estimate-output-series", },
	Short: "Tag Configuration Cardinality Estimator",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.EstimateMetricsOutputSeries(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to estimate-metrics-output-series: %v", err)
		}

		cmdutil.PrintJSON(res, "metric_cardinality_estimate")
	},
}

func init() {
	Cmd.AddCommand(EstimateMetricsOutputSeriesCmd)
}
