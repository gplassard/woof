package usage_metering

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetActiveBillingDimensionsCmd = &cobra.Command{
	Use:   "get_active_billing_dimensions",
	Short: "Get active billing dimensions for cost attribution",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetActiveBillingDimensions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_active_billing_dimensions: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetActiveBillingDimensionsCmd)
}
