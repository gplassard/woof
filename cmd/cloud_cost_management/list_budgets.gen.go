package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListBudgetsCmd = &cobra.Command{
	Use: "list-budgets",

	Short: "List budgets",
	Long: `List budgets
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-budgets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BudgetArray
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.ListBudgets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-budgets")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	Cmd.AddCommand(ListBudgetsCmd)
}
