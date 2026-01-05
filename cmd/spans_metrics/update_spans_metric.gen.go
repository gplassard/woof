package spans_metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSpansMetricCmd = &cobra.Command{
	Use:     "update-spans-metric [metric_id]",
	Aliases: []string{"update"},
	Short:   "Update a span-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateSpansMetric(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SpansMetricUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSpansMetricCmd)
}
