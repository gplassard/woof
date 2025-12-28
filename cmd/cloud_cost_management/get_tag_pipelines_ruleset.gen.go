package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "get_tag_pipelines_ruleset [ruleset_id]",
	Short: "Get a tag pipeline ruleset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.GetTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_tag_pipelines_ruleset: %v", err)
		}

		cmdutil.PrintJSON(res, "ruleset")
	},
}

func init() {
	Cmd.AddCommand(GetTagPipelinesRulesetCmd)
}
