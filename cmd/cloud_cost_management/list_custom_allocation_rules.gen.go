package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCustomAllocationRulesCmd = &cobra.Command{
	Use: "list-custom-allocation-rules",

	Short: "List custom allocation rules",
	Long: `List custom allocation rules
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-custom-allocation-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponseArray
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCustomAllocationRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-custom-allocation-rules")

		fmt.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {

	Cmd.AddCommand(ListCustomAllocationRulesCmd)
}
