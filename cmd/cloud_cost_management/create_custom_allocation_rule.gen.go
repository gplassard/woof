package cloud_cost_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCustomAllocationRuleCmd = &cobra.Command{
	Use: "create-custom-allocation-rule",

	Short: "Create custom allocation rule",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCustomAllocationRule(client.NewContext(apiKey, appKey, site), datadogV2.ArbitraryCostUpsertRequest{})
		cmdutil.HandleError(err, "failed to create-custom-allocation-rule")

		cmd.Println(cmdutil.FormatJSON(res, "arbitrary_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAllocationRuleCmd)
}
