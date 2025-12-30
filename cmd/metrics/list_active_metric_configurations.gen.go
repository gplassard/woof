package metrics

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListActiveMetricConfigurationsCmd = &cobra.Command{
	Use:     "list-active-metric-configurations [metric_name]",
	Aliases: []string{"list-active-configurations"},
	Short:   "List active tags and aggregations",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListActiveMetricConfigurations(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-active-metric-configurations")

		cmd.Println(cmdutil.FormatJSON(res, "actively_queried_configurations"))
	},
}

func init() {
	Cmd.AddCommand(ListActiveMetricConfigurationsCmd)
}
