package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ReorderCustomAllocationRulesCmd = &cobra.Command{
	Use: "reorder-custom-allocation-rules [payload]",

	Short: "Reorder custom allocation rules",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ReorderRuleResourceArray
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err = api.ReorderCustomAllocationRules(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-custom-allocation-rules")

	},
}

func init() {
	Cmd.AddCommand(ReorderCustomAllocationRulesCmd)
}
