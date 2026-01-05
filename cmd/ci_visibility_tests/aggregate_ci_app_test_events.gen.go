package ci_visibility_tests

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AggregateCIAppTestEventsCmd = &cobra.Command{
	Use: "aggregate-ci-app-test-events [payload]",

	Short: "Aggregate tests events",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CIAppTestsAggregateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.AggregateCIAppTestEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-ci-app-test-events")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_tests"))
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppTestEventsCmd)
}
