package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateScorecardRuleCmd = &cobra.Command{
	Use:     "update-scorecard-rule [rule_id]",
	Aliases: []string{"update-rule"},
	Short:   "Update an existing rule",
	Long: `Update an existing rule
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#update-scorecard-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateRuleResponse
		var err error

		var body datadogV2.UpdateRuleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err = api.UpdateScorecardRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-scorecard-rule")

		cmd.Println(cmdutil.FormatJSON(res, "rule"))
	},
}

func init() {

	UpdateScorecardRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateScorecardRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateScorecardRuleCmd)
}
