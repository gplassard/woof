package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateTagPipelinesRulesetCmd = &cobra.Command{
	Use:   "updatetagpipelinesruleset [ruleset_id]",
	Short: "Update tag pipeline ruleset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateRulesetRequest{})
		if err != nil {
			log.Fatalf("failed to updatetagpipelinesruleset: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateTagPipelinesRulesetCmd)
}
