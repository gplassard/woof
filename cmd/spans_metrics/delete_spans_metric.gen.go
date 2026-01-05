package spans_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteSpansMetricCmd = &cobra.Command{
	Use:     "delete-spans-metric [metric_id]",
	Aliases: []string{"delete"},
	Short:   "Delete a span-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		_, err = api.DeleteSpansMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-spans-metric")

	},
}

func init() {
	Cmd.AddCommand(DeleteSpansMetricCmd)
}
