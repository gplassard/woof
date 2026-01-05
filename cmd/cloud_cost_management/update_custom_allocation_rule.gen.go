package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"encoding/json"
	"github.com/spf13/cobra"
	"strconv"
)

var UpdateCustomAllocationRuleCmd = &cobra.Command{
	Use: "update-custom-allocation-rule [rule_id] [payload]",

	Short: "Update custom allocation rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponse
		var err error

		var body datadogV2.ArbitraryCostUpsertRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateCustomAllocationRule(client.NewContext(apiKey, appKey, site), func() int64 { i, _ := strconv.ParseInt(args[0], 10, 64); return i }(), body)
		cmdutil.HandleError(err, "failed to update-custom-allocation-rule")

		cmd.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCustomAllocationRuleCmd)
}
