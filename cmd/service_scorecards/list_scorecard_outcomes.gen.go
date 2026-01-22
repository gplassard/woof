package service_scorecards

import (
	"fmt"
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
	Long: `List all rule outcomes
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#list-scorecard-outcomes`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OutcomesResponse
		var err error

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScorecardOutcomes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scorecard-outcomes")

		fmt.Println(cmdutil.FormatJSON(res, "scorecard_outcome"))
	},
}

func init() {

	Cmd.AddCommand(ListScorecardOutcomesCmd)
}
