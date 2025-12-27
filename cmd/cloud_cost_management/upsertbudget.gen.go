package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpsertBudgetCmd = &cobra.Command{
	Use:   "upsertbudget",
	Short: "Create or update a budget",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpsertBudget(client.NewContext(apiKey, appKey, site), datadogV2.BudgetWithEntries{})
		if err != nil {
			log.Fatalf("failed to upsertbudget: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(UpsertBudgetCmd)
}
