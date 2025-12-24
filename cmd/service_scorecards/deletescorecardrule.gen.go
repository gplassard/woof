package service_scorecards

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteScorecardRuleCmd = &cobra.Command{
	Use:   "deletescorecardrule [rule_id]",
	Short: "Delete a rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err := api.DeleteScorecardRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletescorecardrule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteScorecardRuleCmd)
}
