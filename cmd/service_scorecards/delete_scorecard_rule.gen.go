package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteScorecardRuleCmd = &cobra.Command{
	Use:     "delete-scorecard-rule [rule_id]",
	Aliases: []string{"delete-rule"},
	Short:   "Delete a rule",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err = api.DeleteScorecardRule(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-scorecard-rule")

	},
}

func init() {
	Cmd.AddCommand(DeleteScorecardRuleCmd)
}
