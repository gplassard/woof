package ci_visibility_tests

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchCIAppTestEventsCmd = &cobra.Command{
	Use: "search-ci-app-test-events [payload]",

	Short: "Search tests events",
	Long: `Search tests events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-tests/#search-ci-app-test-events`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppTestEventsResponse
		var err error

		var body datadogV2.SearchCIAppTestEventsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err = api.SearchCIAppTestEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-ci-app-test-events")

		cmd.Println(cmdutil.FormatJSON(res, "citest"))
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppTestEventsCmd)
}
