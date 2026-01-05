package spans_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateSpansMetricCmd = &cobra.Command{
	Use:     "update-spans-metric [metric_id]",
	Aliases: []string{"update"},
	Short:   "Update a span-based metric",
	Long: `Update a span-based metric
Documentation: https://docs.datadoghq.com/api/latest/spans-metrics/#update-spans-metric`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansMetricResponse
		var err error

		var body datadogV2.SpansMetricUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err = api.UpdateSpansMetric(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {

	UpdateSpansMetricCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateSpansMetricCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateSpansMetricCmd)
}
