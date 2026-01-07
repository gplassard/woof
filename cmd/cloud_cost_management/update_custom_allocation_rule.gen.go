package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCustomAllocationRuleCmd = &cobra.Command{
	Use: "update-custom-allocation-rule [rule_id]",

	Short: "Update custom allocation rule",
	Long: `Update custom allocation rule
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#update-custom-allocation-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponse
		var err error

		var body datadogV2.ArbitraryCostUpsertRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCustomAllocationRule(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-custom-allocation-rule")

		fmt.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {

	UpdateCustomAllocationRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCustomAllocationRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCustomAllocationRuleCmd)
}
