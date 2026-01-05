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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceScorecardsApi(client.NewAPIClient())
		_, err := api.UpdateScorecardOutcomesAsync(client.NewContext(apiKey, appKey, site), datadogV2.UpdateOutcomesAsyncRequest{})
		cmdutil.HandleError(err, "failed to update-scorecard-outcomes-async")

	},
}

func init() {
	Cmd.AddCommand(UpdateScorecardOutcomesAsyncCmd)
}
