package metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMetricAssetsCmd = &cobra.Command{
	Use:     "list-metric-assets [metric_name]",
	Aliases: []string{"list-assets"},
	Short:   "Related Assets to a Metric",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListMetricAssets(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-metric-assets")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListMetricAssetsCmd)
}
