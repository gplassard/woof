package usage_metering

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetBillingDimensionMappingCmd = &cobra.Command{
	Use: "get-billing-dimension-mapping",

	Short: "Get billing dimension mapping for usage endpoints",
	Long: `Get billing dimension mapping for usage endpoints
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-billing-dimension-mapping`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BillingDimensionsMappingResponse
		var err error

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetBillingDimensionMapping(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-billing-dimension-mapping")

		fmt.Println(cmdutil.FormatJSON(res, "billing_dimension_mapping"))
	},
}

func init() {

	Cmd.AddCommand(GetBillingDimensionMappingCmd)
}
