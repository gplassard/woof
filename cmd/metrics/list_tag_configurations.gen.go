package metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagConfigurationsCmd = &cobra.Command{
	Use: "list-tag-configurations",

	Short: "Get a list of metrics",
	Long: `Get a list of metrics
Documentation: https://docs.datadoghq.com/api/latest/metrics/#list-tag-configurations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricsAndMetricTagConfigurationsResponse
		var err error

		optionalParams := datadogV2.NewListTagConfigurationsOptionalParameters()

		if cmd.Flags().Changed("filter-configured") {
			val, _ := cmd.Flags().GetString("filter-configured")
			optionalParams.WithFilterConfigured(val)
		}

		if cmd.Flags().Changed("filter-tags-configured") {
			val, _ := cmd.Flags().GetString("filter-tags-configured")
			optionalParams.WithFilterTagsConfigured(val)
		}

		if cmd.Flags().Changed("filter-metric-type") {
			val, _ := cmd.Flags().GetString("filter-metric-type")
			optionalParams.WithFilterMetricType(val)
		}

		if cmd.Flags().Changed("filter-include-percentiles") {
			val, _ := cmd.Flags().GetString("filter-include-percentiles")
			optionalParams.WithFilterIncludePercentiles(val)
		}

		if cmd.Flags().Changed("filter-queried") {
			val, _ := cmd.Flags().GetString("filter-queried")
			optionalParams.WithFilterQueried(val)
		}

		if cmd.Flags().Changed("filter-tags") {
			val, _ := cmd.Flags().GetString("filter-tags")
			optionalParams.WithFilterTags(val)
		}

		if cmd.Flags().Changed("filter-related-assets") {
			val, _ := cmd.Flags().GetString("filter-related-assets")
			optionalParams.WithFilterRelatedAssets(val)
		}

		if cmd.Flags().Changed("window-seconds") {
			val, _ := cmd.Flags().GetInt64("window-seconds")
			optionalParams.WithWindowSeconds(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTagConfigurations(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-tag-configurations")

		fmt.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {

	ListTagConfigurationsCmd.Flags().String("filter-configured", "", "Filter custom metrics that have configured tags.")

	ListTagConfigurationsCmd.Flags().String("filter-tags-configured", "", "Filter tag configurations by configured tags.")

	ListTagConfigurationsCmd.Flags().String("filter-metric-type", "", "Filter metrics by metric type.")

	ListTagConfigurationsCmd.Flags().String("filter-include-percentiles", "", "Filter distributions with additional percentile aggregations enabled or disabled.")

	ListTagConfigurationsCmd.Flags().String("filter-queried", "", "(Preview) Filter custom metrics that have or have not been queried in the specified window[seconds]. If no window is provided or the window is less than 2 hours, a default of 2 hours will be applied.")

	ListTagConfigurationsCmd.Flags().String("filter-tags", "", "Filter metrics that have been submitted with the given tags. Supports boolean and wildcard expressions. Can only be combined with the filter[queried] filter.")

	ListTagConfigurationsCmd.Flags().String("filter-related-assets", "", "(Preview) Filter metrics that are used in dashboards, monitors, notebooks, SLOs.")

	ListTagConfigurationsCmd.Flags().Int64("window-seconds", 0, "The number of seconds of look back (from now) to apply to a filter[tag] or filter[queried] query. Default value is 3600 (1 hour), maximum value is 2,592,000 (30 days).")

	ListTagConfigurationsCmd.Flags().Int64("page-size", 0, "Maximum number of results returned.")

	ListTagConfigurationsCmd.Flags().String("page-cursor", "", "String to query the next page of results. This key is provided with each valid response from the API in 'meta.pagination.next_cursor'. Once the 'meta.pagination.next_cursor' key is null, all pages have been retrieved.")

	Cmd.AddCommand(ListTagConfigurationsCmd)
}
