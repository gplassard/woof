package rum_metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRumMetricCmd = &cobra.Command{
	Use:   "get-rum-metric [metric_id]",
	Aliases: []string{ "get", },
	Short: "Get a rum-based metric",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.GetRumMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-rum-metric")

		cmdutil.PrintJSON(res, "rum_metrics")
	},
}

func init() {
	Cmd.AddCommand(GetRumMetricCmd)
}
