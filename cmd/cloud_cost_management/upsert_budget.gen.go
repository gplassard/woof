package cloud_cost_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertBudgetCmd = &cobra.Command{
	Use: "upsert-budget",

	Short: "Create or update a budget",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpsertBudget(client.NewContext(apiKey, appKey, site), datadogV2.BudgetWithEntries{})
		cmdutil.HandleError(err, "failed to upsert-budget")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {
	Cmd.AddCommand(UpsertBudgetCmd)
}
