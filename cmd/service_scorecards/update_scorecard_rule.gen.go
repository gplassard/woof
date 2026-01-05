package service_scorecards

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateScorecardRuleCmd = &cobra.Command{
	Use:     "update-scorecard-rule [rule_id]",
	Aliases: []string{"update-rule"},
	Short:   "Update an existing rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.UpdateScorecardRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateRuleRequest{})
		cmdutil.HandleError(err, "failed to update-scorecard-rule")

		cmd.Println(cmdutil.FormatJSON(res, "rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardRuleCmd)
}
