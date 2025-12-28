package synthetics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOnDemandConcurrencyCapCmd = &cobra.Command{
	Use:   "get_on_demand_concurrency_cap",
	Short: "Get the on-demand concurrency cap",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		res, _, err := api.GetOnDemandConcurrencyCap(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_on_demand_concurrency_cap: %v", err)
		}

		cmdutil.PrintJSON(res, "on_demand_concurrency_cap")
	},
}

func init() {
	Cmd.AddCommand(GetOnDemandConcurrencyCapCmd)
}
