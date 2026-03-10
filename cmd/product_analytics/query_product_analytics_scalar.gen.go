package product_analytics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var QueryProductAnalyticsScalarCmd = &cobra.Command{
	Use:     "query-product-analytics-scalar",
	Aliases: []string{"query-scalar"},
	Short:   "Compute scalar analytics",
	Long: `Compute scalar analytics
Documentation: https://docs.datadoghq.com/api/latest/product-analytics/#query-product-analytics-scalar`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProductAnalyticsScalarResponse
		var err error

		var body datadogV2.ProductAnalyticsAnalyticsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewProductAnalyticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.QueryProductAnalyticsScalar(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to query-product-analytics-scalar")

		fmt.Println(cmdutil.FormatJSON(res, "scalar_response"))
	},
}

func init() {

	QueryProductAnalyticsScalarCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	QueryProductAnalyticsScalarCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(QueryProductAnalyticsScalarCmd)
}
