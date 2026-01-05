package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateScorecardOutcomesBatchCmd = &cobra.Command{
	Use:     "create-scorecard-outcomes-batch [payload]",
	Aliases: []string{"create-outcomes-batch"},
	Short:   "Create outcomes batch",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OutcomesBatchResponse
		var err error

		var body datadogV2.OutcomesBatchRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err = api.CreateScorecardOutcomesBatch(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scorecard-outcomes-batch")

		cmd.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {
	Cmd.AddCommand(CreateScorecardOutcomesBatchCmd)
}
