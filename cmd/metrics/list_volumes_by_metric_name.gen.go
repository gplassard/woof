package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListVolumesByMetricNameCmd = &cobra.Command{
	Use:   "list-volumes-by-metric-name [metric_name]",
	Short: "List distinct metric volumes by metric name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListVolumesByMetricName(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-volumes-by-metric-name: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(ListVolumesByMetricNameCmd)
}
