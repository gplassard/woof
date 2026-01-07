package spans_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSpansMetricCmd = &cobra.Command{
	Use:     "get-spans-metric [metric_id]",
	Aliases: []string{"get"},
	Short:   "Get a span-based metric",
	Long: `Get a span-based metric
Documentation: https://docs.datadoghq.com/api/latest/spans-metrics/#get-spans-metric`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansMetricResponse
		var err error

		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSpansMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-spans-metric")

		fmt.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {

	Cmd.AddCommand(GetSpansMetricCmd)
}
