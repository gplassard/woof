package service_scorecards

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListScorecardOutcomesCmd = &cobra.Command{
	Use:   "list-scorecard-outcomes",
	Short: "List all rule outcomes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.ListScorecardOutcomes(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-scorecard-outcomes: %v", err)
		}

		cmdutil.PrintJSON(res, "service_scorecards")
	},
}

func init() {
	Cmd.AddCommand(ListScorecardOutcomesCmd)
}
