package service_scorecards

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateScorecardRuleCmd = &cobra.Command{
	Use:   "createscorecardrule",
	Short: "Create a new rule",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.CreateScorecardRule(client.NewContext(apiKey, appKey, site), datadogV2.CreateRuleRequest{})
		if err != nil {
			log.Fatalf("failed to createscorecardrule: %v", err)
		}

		cmdutil.PrintJSON(res, "service_scorecards")
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardRuleCmd)
}
