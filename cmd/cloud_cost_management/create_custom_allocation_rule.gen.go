package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCustomAllocationRuleCmd = &cobra.Command{
	Use: "create-custom-allocation-rule [payload]",

	Short: "Create custom allocation rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ArbitraryRuleResponse
		var err error

		var body datadogV2.ArbitraryCostUpsertRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateCustomAllocationRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-custom-allocation-rule")

		cmd.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAllocationRuleCmd)
}
