package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTagPipelinesRulesetsCmd = &cobra.Command{
	Use:   "listtagpipelinesrulesets",
	Short: "List tag pipeline rulesets",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListTagPipelinesRulesets(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listtagpipelinesrulesets: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(ListTagPipelinesRulesetsCmd)
}
