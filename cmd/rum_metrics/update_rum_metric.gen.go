package rum_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRumMetricCmd = &cobra.Command{
	Use:     "update-rum-metric [metric_id]",
	Aliases: []string{"update"},
	Short:   "Update a rum-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateRumMetric(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RumMetricUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-rum-metric")

		cmd.Println(cmdutil.FormatJSON(res, "rum_metrics"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRumMetricCmd)
}
