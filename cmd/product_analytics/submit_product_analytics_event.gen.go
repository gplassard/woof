package product_analytics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SubmitProductAnalyticsEventCmd = &cobra.Command{
	Use:     "submit-product-analytics-event",
	Aliases: []string{"submit-event"},
	Short:   "Send server-side events",
	Long: `Send server-side events
Documentation: https://docs.datadoghq.com/api/latest/product-analytics/#submit-product-analytics-event`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		var body datadogV2.ProductAnalyticsServerSideEventItem
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewProductAnalyticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SubmitProductAnalyticsEvent(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to submit-product-analytics-event")

		fmt.Println(cmdutil.FormatJSON(res, "product_analytics"))
	},
}

func init() {

	SubmitProductAnalyticsEventCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SubmitProductAnalyticsEventCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SubmitProductAnalyticsEventCmd)
}
