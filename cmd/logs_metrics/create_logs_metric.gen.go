package logs_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLogsMetricCmd = &cobra.Command{
	Use:     "create-logs-metric",
	Aliases: []string{"create"},
	Short:   "Create a log-based metric",
	Long: `Create a log-based metric
Documentation: https://docs.datadoghq.com/api/latest/logs-metrics/#create-logs-metric`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsMetricResponse
		var err error

		var body datadogV2.LogsMetricCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLogsMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateLogsMetric(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-logs-metric")

		fmt.Println(cmdutil.FormatJSON(res, "logs_metric"))
	},
}

func init() {

	CreateLogsMetricCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLogsMetricCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLogsMetricCmd)
}
