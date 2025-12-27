package service_scorecards

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateScorecardOutcomesBatchCmd = &cobra.Command{
	Use:   "create_scorecard_outcomes_batch",
	Short: "Create outcomes batch",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.CreateScorecardOutcomesBatch(client.NewContext(apiKey, appKey, site), datadogV2.OutcomesBatchRequest{})
		if err != nil {
			log.Fatalf("failed to create_scorecard_outcomes_batch: %v", err)
		}

		cmdutil.PrintJSON(res, "service_scorecards")
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardOutcomesBatchCmd)
}
