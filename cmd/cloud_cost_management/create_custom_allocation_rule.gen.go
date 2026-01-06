package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCustomAllocationRuleCmd = &cobra.Command{
	Use: "create-custom-allocation-rule",

	Short: "Create custom allocation rule",
	Long: `Create custom allocation rule
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-custom-allocation-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponse
		var err error

		var body datadogV2.ArbitraryCostUpsertRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCustomAllocationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-custom-allocation-rule")

		cmd.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {

	CreateCustomAllocationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCustomAllocationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCustomAllocationRuleCmd)
}
