package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReorderCustomAllocationRulesCmd = &cobra.Command{
	Use:   "reordercustomallocationrules",
	Short: "Reorder custom allocation rules",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.ReorderCustomAllocationRules(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRuleResourceArray{})
		if err != nil {
			log.Fatalf("failed to reordercustomallocationrules: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ReorderCustomAllocationRulesCmd)
}
