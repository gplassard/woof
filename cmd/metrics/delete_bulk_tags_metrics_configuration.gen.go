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
	Long: `Delete tags for multiple metrics
Documentation: https://docs.datadoghq.com/api/latest/metrics/#delete-bulk-tags-metrics-configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricBulkTagConfigResponse
		var err error

		var body datadogV2.MetricBulkTagConfigDeleteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DeleteBulkTagsMetricsConfiguration(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-bulk-tags-metrics-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "metric_bulk_configure_tags"))
	},
}

func init() {

	DeleteBulkTagsMetricsConfigurationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteBulkTagsMetricsConfigurationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteBulkTagsMetricsConfigurationCmd)
}
