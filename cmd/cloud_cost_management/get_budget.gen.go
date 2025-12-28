package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetBudgetCmd = &cobra.Command{
	Use:   "get-budget [budget_id]",
	
	Short: "Get a budget",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.GetBudget(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-budget: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(GetBudgetCmd)
}
