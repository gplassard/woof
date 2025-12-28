package service_scorecards

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteScorecardRuleCmd = &cobra.Command{
	Use:   "delete-scorecard-rule [rule_id]",
	Short: "Delete a rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err := api.DeleteScorecardRule(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-scorecard-rule: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteScorecardRuleCmd)
}
