package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "delete_tag_pipelines_ruleset [ruleset_id]",
	Short: "Delete tag pipeline ruleset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.DeleteTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_tag_pipelines_ruleset: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteTagPipelinesRulesetCmd)
}
