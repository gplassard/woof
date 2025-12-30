package metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SubmitMetricsCmd = &cobra.Command{
	Use:     "submit-metrics",
	Aliases: []string{"submit"},
	Short:   "Submit metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.SubmitMetrics(client.NewContext(apiKey, appKey, site), datadogV2.MetricPayload{})
		cmdutil.HandleError(err, "failed to submit-metrics")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(SubmitMetricsCmd)
}
