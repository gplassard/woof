package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "create-tag-pipelines-ruleset",
	Short: "Create tag pipeline ruleset",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.CreateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), datadogV2.CreateRulesetRequest{})
		if err != nil {
			log.Fatalf("failed to create-tag-pipelines-ruleset: %v", err)
		}

		cmdutil.PrintJSON(res, "ruleset")
	},
}

func init() {
	Cmd.AddCommand(CreateTagPipelinesRulesetCmd)
}
