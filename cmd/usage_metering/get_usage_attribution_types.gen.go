package usage_metering

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUsageAttributionTypesCmd = &cobra.Command{
	Use: "get-usage-attribution-types",

	Short: "Get usage attribution types",
	Long: `Get usage attribution types
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-usage-attribution-types`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsageAttributionTypesResponse
		var err error

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetUsageAttributionTypes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-usage-attribution-types")

		fmt.Println(cmdutil.FormatJSON(res, "usage_attribution_types"))
	},
}

func init() {

	Cmd.AddCommand(GetUsageAttributionTypesCmd)
}
