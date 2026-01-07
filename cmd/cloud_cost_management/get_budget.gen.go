package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetBudgetCmd = &cobra.Command{
	Use: "get-budget [budget_id]",

	Short: "Get a budget",
	Long: `Get a budget
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#get-budget`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BudgetWithEntries
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetBudget(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-budget")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	Cmd.AddCommand(GetBudgetCmd)
}
