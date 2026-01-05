package usage_metering

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetActiveBillingDimensionsCmd = &cobra.Command{
	Use: "get-active-billing-dimensions",

	Short: "Get active billing dimensions for cost attribution",
	Long: `Get active billing dimensions for cost attribution
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-active-billing-dimensions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ActiveBillingDimensionsResponse
		var err error

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err = api.GetActiveBillingDimensions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-active-billing-dimensions")

		cmd.Println(cmdutil.FormatJSON(res, "billing_dimensions"))
	},
}

func init() {
	Cmd.AddCommand(GetActiveBillingDimensionsCmd)
}
