package usage_metering

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetProjectedCostCmd = &cobra.Command{
	Use:   "get-projected-cost",
	
	Short: "Get projected cost across your account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetProjectedCost(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-projected-cost: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetProjectedCostCmd)
}
