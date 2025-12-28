package service_scorecards

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateScorecardRuleCmd = &cobra.Command{
	Use:   "create-scorecard-rule",
	
	Short: "Create a new rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.CreateScorecardRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateRuleRequest{})
		cmdutil.HandleError(err, "failed to create-scorecard-rule")

		cmdutil.PrintJSON(res, "rule")
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardRuleCmd)
}
