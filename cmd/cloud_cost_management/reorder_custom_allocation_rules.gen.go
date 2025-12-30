package cloud_cost_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderCustomAllocationRulesCmd = &cobra.Command{
	Use: "reorder-custom-allocation-rules",

	Short: "Reorder custom allocation rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.ReorderCustomAllocationRules(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRuleResourceArray{})
		cmdutil.HandleError(err, "failed to reorder-custom-allocation-rules")

	},
}

func init() {
	Cmd.AddCommand(ReorderCustomAllocationRulesCmd)
}
