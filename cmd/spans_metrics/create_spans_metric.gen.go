package spans_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSpansMetricCmd = &cobra.Command{
	Use:     "create-spans-metric [payload]",
	Aliases: []string{"create"},
	Short:   "Create a span-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SpansMetricCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSpansMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateSpansMetric(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-spans-metric")

		cmd.Println(cmdutil.FormatJSON(res, "spans_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateSpansMetricCmd)
}
