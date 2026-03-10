package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateBudgetCmd = &cobra.Command{
	Use: "validate-budget",

	Short: "Validate budget",
	Long: `Validate budget
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#validate-budget`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BudgetValidationResponse
		var err error

		var body datadogV2.BudgetValidationRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ValidateBudget(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-budget")

		fmt.Println(cmdutil.FormatJSON(res, "budget_validation"))
	},
}

func init() {

	ValidateBudgetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateBudgetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateBudgetCmd)
}
