package usage_metering

import (
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

		optionalParams := datadogV2.NewGetBillingDimensionMappingOptionalParameters()

		if cmd.Flags().Changed("filter-month") {
			val, _ := cmd.Flags().GetString("filter-month")
			optionalParams.WithFilterMonth(val)
		}

		if cmd.Flags().Changed("filter-view") {
			val, _ := cmd.Flags().GetString("filter-view")
			optionalParams.WithFilterView(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetBillingDimensionMapping(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-billing-dimension-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {

	GetBillingDimensionMappingCmd.Flags().String("filter-month", "", "Datetime in ISO-8601 format, UTC, and for mappings beginning this month. Defaults to the current month.")

	GetBillingDimensionMappingCmd.Flags().String("filter-view", "", "String to specify whether to retrieve active billing dimension mappings for the contract or for all available mappings. Allowed views have the string 'active' or 'all'. Defaults to 'active'.")

	Cmd.AddCommand(GetBillingDimensionMappingCmd)
}
