package usage_metering

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetEstimatedCostByOrgCmd = &cobra.Command{
	Use:   "get_estimated_cost_by_org",
	Short: "Get estimated cost across your account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		res, _, err := api.GetEstimatedCostByOrg(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_estimated_cost_by_org: %v", err)
		}

		cmdutil.PrintJSON(res, "usage_metering")
	},
}

func init() {
	Cmd.AddCommand(GetEstimatedCostByOrgCmd)
}
