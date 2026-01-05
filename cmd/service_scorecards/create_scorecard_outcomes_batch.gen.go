package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateScorecardOutcomesBatchCmd = &cobra.Command{
	Use:     "create-scorecard-outcomes-batch",
	Aliases: []string{"create-outcomes-batch"},
	Short:   "Create outcomes batch",
	Long: `Create outcomes batch
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#create-scorecard-outcomes-batch`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OutcomesBatchResponse
		var err error

		var body datadogV2.OutcomesBatchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		res, _, err = api.CreateScorecardOutcomesBatch(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-scorecard-outcomes-batch")

		cmd.Println(cmdutil.FormatJSON(res, "service_scorecards"))
	},
}

func init() {

	CreateScorecardOutcomesBatchCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateScorecardOutcomesBatchCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateScorecardOutcomesBatchCmd)
}
