package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReorderTagPipelinesRulesetsCmd = &cobra.Command{
	Use:   "reorder-tag-pipelines-rulesets",
	Short: "Reorder tag pipeline rulesets",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.ReorderTagPipelinesRulesets(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRulesetResourceArray{})
		if err != nil {
			log.Fatalf("failed to reorder-tag-pipelines-rulesets: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ReorderTagPipelinesRulesetsCmd)
}
