package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpsertBudgetCmd = &cobra.Command{
	Use: "upsert-budget [payload]",

	Short: "Create or update a budget",
	Long: `Create or update a budget
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#upsert-budget`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BudgetWithEntries
		var err error

		var body datadogV2.BudgetWithEntries
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.UpsertBudget(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to upsert-budget")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {
	Cmd.AddCommand(UpsertBudgetCmd)
}
