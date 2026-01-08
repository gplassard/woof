package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderCustomAllocationRulesCmd = &cobra.Command{
	Use: "reorder-custom-allocation-rules",

	Short: "Reorder custom allocation rules",
	Long: `Reorder custom allocation rules
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#reorder-custom-allocation-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ReorderRuleResourceArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ReorderCustomAllocationRules(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-custom-allocation-rules")

	},
}

func init() {

	ReorderCustomAllocationRulesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ReorderCustomAllocationRulesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ReorderCustomAllocationRulesCmd)
}
