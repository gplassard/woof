package product_analytics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryProductAnalyticsTimeseriesCmd = &cobra.Command{
	Use:     "query-product-analytics-timeseries",
	Aliases: []string{"query-timeseries"},
	Short:   "Compute timeseries analytics",
	Long: `Compute timeseries analytics
Documentation: https://docs.datadoghq.com/api/latest/product-analytics/#query-product-analytics-timeseries`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProductAnalyticsTimeseriesResponse
		var err error

		var body datadogV2.ProductAnalyticsAnalyticsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewProductAnalyticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.QueryProductAnalyticsTimeseries(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-product-analytics-timeseries")

		fmt.Println(cmdutil.FormatJSON(res, "timeseries_response"))
	},
}

func init() {

	QueryProductAnalyticsTimeseriesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryProductAnalyticsTimeseriesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryProductAnalyticsTimeseriesCmd)
}
