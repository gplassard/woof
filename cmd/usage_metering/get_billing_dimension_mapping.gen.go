package usage_metering

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetBillingDimensionMappingCmd = &cobra.Command{
	Use: "get-billing-dimension-mapping",

	Short: "Get billing dimension mapping for usage endpoints",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetBillingDimensionMapping(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-billing-dimension-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "usage_metering"))
	},
}

func init() {
	Cmd.AddCommand(GetBillingDimensionMappingCmd)
}
