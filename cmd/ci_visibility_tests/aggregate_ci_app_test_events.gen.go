package ci_visibility_tests

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateCIAppTestEventsCmd = &cobra.Command{
	Use:   "aggregate-ci-app-test-events",
	
	Short: "Aggregate tests events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.AggregateCIAppTestEvents(client.NewContext(apiKey, appKey, site), datadogV2.CIAppTestsAggregateRequest{})
		cmdutil.HandleError(err, "failed to aggregate-ci-app-test-events")

		cmdutil.PrintJSON(res, "ci_visibility_tests")
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppTestEventsCmd)
}
