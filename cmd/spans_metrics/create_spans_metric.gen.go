package spans_metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSpansMetricCmd = &cobra.Command{
	Use:     "create-spans-metric",
	Aliases: []string{"create"},
	Short:   "Create a span-based metric",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateSpansMetric(client.NewContext(apiKey, appKey, site), datadogV2.SpansMetricCreateRequest{})
		cmdutil.HandleError(err, "failed to create-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateSpansMetricCmd)
}
