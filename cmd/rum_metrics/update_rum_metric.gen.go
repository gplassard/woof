package rum_metrics

import (
	"fmt"
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
	Long: `Update a rum-based metric
Documentation: https://docs.datadoghq.com/api/latest/rum-metrics/#update-rum-metric`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumMetricResponse
		var err error

		var body datadogV2.RumMetricUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateRumMetric(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-rum-metric")

		fmt.Println(cmdutil.FormatJSON(res, "rum_metric"))
	},
}

func init() {

	UpdateRumMetricCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRumMetricCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRumMetricCmd)
}
