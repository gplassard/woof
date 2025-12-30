package service_scorecards

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScorecardRulesCmd = &cobra.Command{
	Use:     "list-scorecard-rules",
	Aliases: []string{"list-rules"},
	Short:   "List all rules",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.ListScorecardRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scorecard-rules")

		cmd.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {
	Cmd.AddCommand(ListScorecardRulesCmd)
}
