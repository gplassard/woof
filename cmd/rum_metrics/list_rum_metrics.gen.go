package rum_metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRumMetricsCmd = &cobra.Command{
	Use:     "list-rum-metrics",
	Aliases: []string{"list"},
	Short:   "Get all rum-based metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumMetricsApi(client.NewAPIClient())
		res, _, err := api.ListRumMetrics(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-rum-metrics")

		cmd.Println(cmdutil.FormatJSON(res, "rum_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListRumMetricsCmd)
}
