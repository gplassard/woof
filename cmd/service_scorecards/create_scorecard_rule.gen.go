package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateScorecardRuleCmd = &cobra.Command{
	Use:     "create-scorecard-rule [payload]",
	Aliases: []string{"create-rule"},
	Short:   "Create a new rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CreateRuleRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.CreateScorecardRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scorecard-rule")

		cmd.Println(cmdutil.FormatJSON(res, "rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardRuleCmd)
}
