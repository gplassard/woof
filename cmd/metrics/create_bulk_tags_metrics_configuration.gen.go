package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:     "create-bulk-tags-metrics-configuration",
	Aliases: []string{"create-bulk-tags-configuration"},
	Short:   "Configure tags for multiple metrics",
	Long: `Configure tags for multiple metrics
Documentation: https://docs.datadoghq.com/api/latest/metrics/#create-bulk-tags-metrics-configuration`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricBulkTagConfigResponse
		var err error

		var body datadogV2.MetricBulkTagConfigCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateBulkTagsMetricsConfiguration(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-bulk-tags-metrics-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "metric_bulk_configure_tags"))
	},
}

func init() {

	CreateBulkTagsMetricsConfigurationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateBulkTagsMetricsConfigurationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateBulkTagsMetricsConfigurationCmd)
}
