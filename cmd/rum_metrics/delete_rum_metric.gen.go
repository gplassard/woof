package rum_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRumMetricCmd = &cobra.Command{
	Use:     "delete-rum-metric [metric_id]",
	Aliases: []string{"delete"},
	Short:   "Delete a rum-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		_, err := api.DeleteRumMetric(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-rum-metric")

	},
}

func init() {
	Cmd.AddCommand(DeleteRumMetricCmd)
}
