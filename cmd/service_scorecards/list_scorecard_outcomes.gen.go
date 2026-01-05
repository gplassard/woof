package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScorecardOutcomesCmd = &cobra.Command{
	Use:     "list-scorecard-outcomes",
	Aliases: []string{"list-outcomes"},
	Short:   "List all rule outcomes",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.ListScorecardOutcomes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scorecard-outcomes")

		cmd.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {
	Cmd.AddCommand(ListScorecardOutcomesCmd)
}
