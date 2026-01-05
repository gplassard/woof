package spans_metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSpansMetricCmd = &cobra.Command{
	Use:     "get-spans-metric [metric_id]",
	Aliases: []string{"get"},
	Short:   "Get a span-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.GetSpansMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(GetSpansMetricCmd)
}
