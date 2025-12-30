package rum_metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRumMetricCmd = &cobra.Command{
	Use:     "create-rum-metric",
	Aliases: []string{"create"},
	Short:   "Create a rum-based metric",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateRumMetric(client.NewContext(apiKey, appKey, site), datadogV2.RumMetricCreateRequest{})
		cmdutil.HandleError(err, "failed to create-rum-metric")

		cmd.Println(cmdutil.FormatJSON(res, "rum_metrics"))
	},
}

func init() {
	Cmd.AddCommand(CreateRumMetricCmd)
}
