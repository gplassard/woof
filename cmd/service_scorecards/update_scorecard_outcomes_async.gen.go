package service_scorecards

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateScorecardOutcomesAsyncCmd = &cobra.Command{
	Use:     "update-scorecard-outcomes-async",
	Aliases: []string{"update-outcomes-async"},
	Short:   "Update Scorecard outcomes asynchronously",
	Long: `Update Scorecard outcomes asynchronously
Documentation: https://docs.datadoghq.com/api/latest/service-scorecards/#update-scorecard-outcomes-async`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.UpdateOutcomesAsyncRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err = api.UpdateScorecardOutcomesAsync(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-scorecard-outcomes-async")

	},
}

func init() {

	UpdateScorecardOutcomesAsyncCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateScorecardOutcomesAsyncCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateScorecardOutcomesAsyncCmd)
}
