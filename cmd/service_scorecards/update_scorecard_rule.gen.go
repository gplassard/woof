package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateScorecardRuleCmd = &cobra.Command{
	Use:     "update-scorecard-rule [rule_id] [payload]",
	Aliases: []string{"update-rule"},
	Short:   "Update an existing rule",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UpdateRuleRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.UpdateScorecardRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-scorecard-rule")

		cmd.Println(cmdutil.FormatJSON(res, "rule"))
	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardRuleCmd)
}
