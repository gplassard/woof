package service_scorecards

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScorecardOutcomesBatchCmd = &cobra.Command{
	Use:     "create-scorecard-outcomes-batch",
	Aliases: []string{"create-outcomes-batch"},
	Short:   "Create outcomes batch",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err := api.CreateScorecardOutcomesBatch(client.NewContext(apiKey, appKey, site), datadogV2.OutcomesBatchRequest{})
		cmdutil.HandleError(err, "failed to create-scorecard-outcomes-batch")

		cmd.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardOutcomesBatchCmd)
}
