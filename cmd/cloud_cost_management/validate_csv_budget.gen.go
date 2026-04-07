package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateCsvBudgetCmd = &cobra.Command{
	Use: "validate-csv-budget",

	Short: "ValidateCsvBudget",
	Long: `ValidateCsvBudget
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#validate-csv-budget`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ValidationResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ValidateCsvBudget(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to validate-csv-budget")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	Cmd.AddCommand(ValidateCsvBudgetCmd)
}
