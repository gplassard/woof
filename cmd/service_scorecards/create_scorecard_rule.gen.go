package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScorecardRuleCmd = &cobra.Command{
	Use:     "create-scorecard-rule",
	Aliases: []string{"create-rule"},
	Short:   "Create a new rule",
	Long: `Create a new rule
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#create-scorecard-rule`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateRuleResponse
		var err error

		var body datadogV2.CreateRuleRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateScorecardRule(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scorecard-rule")

		cmd.Println(cmdutil.FormatJSON(res, "rule"))
	},
}

func init() {

	CreateScorecardRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateScorecardRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateScorecardRuleCmd)
}
