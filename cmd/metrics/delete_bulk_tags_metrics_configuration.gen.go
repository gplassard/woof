package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:     "delete-bulk-tags-metrics-configuration",
	Aliases: []string{"delete-bulk-tags-configuration"},
	Short:   "Delete tags for multiple metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.DeleteBulkTagsMetricsConfiguration(client.NewContext(apiKey, appKey, site), datadogV2.MetricBulkTagConfigDeleteRequest{})
		cmdutil.HandleError(err, "failed to delete-bulk-tags-metrics-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "metric_bulk_configure_tags"))
	},
}

func init() {
	Cmd.AddCommand(DeleteBulkTagsMetricsConfigurationCmd)
}
