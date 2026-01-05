package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagsByMetricNameCmd = &cobra.Command{
	Use:     "list-tags-by-metric-name [metric_name]",
	Aliases: []string{"list-tags-by-name"},
	Short:   "List tags by metric name",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagsByMetricName(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-tags-by-metric-name")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListTagsByMetricNameCmd)
}
