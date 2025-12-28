package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCustomAllocationRuleCmd = &cobra.Command{
	Use:   "create_custom_allocation_rule",
	Short: "Create custom allocation rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCustomAllocationRule(client.NewContext(apiKey, appKey, site), datadogV2.ArbitraryCostUpsertRequest{})
		if err != nil {
			log.Fatalf("failed to create_custom_allocation_rule: %v", err)
		}

		cmdutil.PrintJSON(res, "arbitrary_rule")
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAllocationRuleCmd)
}
