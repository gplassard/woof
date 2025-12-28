package service_scorecards

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateScorecardOutcomesAsyncCmd = &cobra.Command{
	Use:   "update-scorecard-outcomes-async",
	Short: "Update Scorecard outcomes asynchronously",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err := api.UpdateScorecardOutcomesAsync(client.NewContext(apiKey, appKey, site), datadogV2.UpdateOutcomesAsyncRequest{})
		if err != nil {
			log.Fatalf("failed to update-scorecard-outcomes-async: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardOutcomesAsyncCmd)
}
