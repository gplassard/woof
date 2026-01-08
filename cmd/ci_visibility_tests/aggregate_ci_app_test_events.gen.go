package ci_visibility_tests

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AggregateCIAppTestEventsCmd = &cobra.Command{
	Use: "aggregate-ci-app-test-events",

	Short: "Aggregate tests events",
	Long: `Aggregate tests events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-tests/#aggregate-ci-app-test-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppTestsAnalyticsAggregateResponse
		var err error

		var body datadogV2.CIAppTestsAggregateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AggregateCIAppTestEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-ci-app-test-events")

		fmt.Println(cmdutil.FormatJSON(res, "ci_visibility_tests"))
	},
}

func init() {

	AggregateCIAppTestEventsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AggregateCIAppTestEventsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AggregateCIAppTestEventsCmd)
}
