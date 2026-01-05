package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SubmitMetricsCmd = &cobra.Command{
	Use:     "submit-metrics [payload]",
	Aliases: []string{"submit"},
	Short:   "Submit metrics",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.MetricPayload
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.SubmitMetrics(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to submit-metrics")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(SubmitMetricsCmd)
}
