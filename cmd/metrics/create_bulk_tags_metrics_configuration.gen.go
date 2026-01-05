package metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:     "create-bulk-tags-metrics-configuration",
	Aliases: []string{"create-bulk-tags-configuration"},
	Short:   "Configure tags for multiple metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateBulkTagsMetricsConfiguration(client.NewContext(apiKey, appKey, site), datadogV2.MetricBulkTagConfigCreateRequest{})
		cmdutil.HandleError(err, "failed to create-bulk-tags-metrics-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "metric_bulk_configure_tags"))
	},
}

func init() {
	Cmd.AddCommand(CreateBulkTagsMetricsConfigurationCmd)
}
