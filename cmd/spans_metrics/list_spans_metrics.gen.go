package spans_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSpansMetricsCmd = &cobra.Command{
	Use:     "list-spans-metrics",
	Aliases: []string{"list"},
	Short:   "Get all span-based metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.ListSpansMetrics(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-spans-metrics")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListSpansMetricsCmd)
}
