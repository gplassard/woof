package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTagsByMetricNameCmd = &cobra.Command{
	Use:   "list-tags-by-metric-name [metric_name]",
	Aliases: []string{ "list-tags-by-name", },
	Short: "List tags by metric name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagsByMetricName(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-tags-by-metric-name: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(ListTagsByMetricNameCmd)
}
