package service_scorecards

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListScorecardRulesCmd = &cobra.Command{
	Use:     "list-scorecard-rules",
	Aliases: []string{"list-rules"},
	Short:   "List all rules",
	Long: `List all rules
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#list-scorecard-rules`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListRulesResponse
		var err error

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListScorecardRules(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-scorecard-rules")

		fmt.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {

	Cmd.AddCommand(ListScorecardRulesCmd)
}
