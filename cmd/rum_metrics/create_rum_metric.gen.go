package rum_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRumMetricCmd = &cobra.Command{
	Use:     "create-rum-metric",
	Aliases: []string{"create"},
	Short:   "Create a rum-based metric",
	Long: `Create a rum-based metric
Documentation: https://docs.datadoghq.com/api/latest/rum-metrics/#create-rum-metric`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumMetricResponse
		var err error

		var body datadogV2.RumMetricCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateRumMetric(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-rum-metric")

		fmt.Println(cmdutil.FormatJSON(res, "rum_metric"))
	},
}

func init() {

	CreateRumMetricCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateRumMetricCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateRumMetricCmd)
}
