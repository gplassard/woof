package service_scorecards

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateScorecardRuleCmd = &cobra.Command{
	Use:   "update-scorecard-rule [rule_id]",
	Short: "Update an existing rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.UpdateScorecardRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateRuleRequest{})
		if err != nil {
			log.Fatalf("failed to update-scorecard-rule: %v", err)
		}

		cmdutil.PrintJSON(res, "rule")
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardRuleCmd)
}
