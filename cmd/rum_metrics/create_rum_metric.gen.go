package rum_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateRumMetricCmd = &cobra.Command{
	Use:     "create-rum-metric [payload]",
	Aliases: []string{"create"},
	Short:   "Create a rum-based metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumMetricResponse
		var err error

		var body datadogV2.RumMetricCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err = api.CreateRumMetric(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-rum-metric")

		cmd.Println(cmdutil.FormatJSON(res, "rum_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateRumMetricCmd)
}
