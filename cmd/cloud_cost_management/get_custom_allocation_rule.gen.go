package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var GetCustomAllocationRuleCmd = &cobra.Command{
	Use: "get-custom-allocation-rule [rule_id]",

	Short: "Get custom allocation rule",
	Long: `Get custom allocation rule
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#get-custom-allocation-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.GetCustomAllocationRule(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }())
		cmdutil.HandleError(err, "failed to get-custom-allocation-rule")

		cmd.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomAllocationRuleCmd)
}
