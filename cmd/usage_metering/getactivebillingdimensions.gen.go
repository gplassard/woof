package usage_metering

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetActiveBillingDimensionsCmd = &cobra.Command{
	Use:   "getactivebillingdimensions",
	Short: "Get active billing dimensions for cost attribution",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetActiveBillingDimensions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getactivebillingdimensions: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetActiveBillingDimensionsCmd)
}
