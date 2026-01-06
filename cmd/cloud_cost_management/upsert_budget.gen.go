package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertBudgetCmd = &cobra.Command{
	Use: "upsert-budget",

	Short: "Create or update a budget",
	Long: `Create or update a budget
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#upsert-budget`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BudgetWithEntries
		var err error

		var body datadogV2.BudgetWithEntries
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpsertBudget(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to upsert-budget")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	UpsertBudgetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpsertBudgetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpsertBudgetCmd)
}
