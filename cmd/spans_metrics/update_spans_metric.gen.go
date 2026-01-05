package spans_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateSpansMetricCmd = &cobra.Command{
	Use:     "update-spans-metric [metric_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update a span-based metric",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansMetricResponse
		var err error

		var body datadogV2.SpansMetricUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err = api.UpdateSpansMetric(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(UpdateSpansMetricCmd)
}
