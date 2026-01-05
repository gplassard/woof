package cloud_cost_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

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
