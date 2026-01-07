package metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SubmitMetricsCmd = &cobra.Command{
	Use:     "submit-metrics",
	Aliases: []string{"submit"},
	Short:   "Submit metrics",
	Long: `Submit metrics
Documentation: https://docs.datadoghq.com/api/latest/metrics/#submit-metrics`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IntakePayloadAccepted
		var err error

		var body datadogV2.MetricPayload
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SubmitMetrics(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to submit-metrics")

		fmt.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {

	SubmitMetricsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SubmitMetricsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SubmitMetricsCmd)
}
