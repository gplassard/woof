package ci_visibility_tests

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchCIAppTestEventsCmd = &cobra.Command{
	Use: "search-ci-app-test-events",

	Short: "Search tests events",
	Long: `Search tests events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-tests/#search-ci-app-test-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppTestEventsResponse
		var err error

		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchCIAppTestEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to search-ci-app-test-events")

		fmt.Println(cmdutil.FormatJSON(res, "c_i_app_test_event"))
	},
}

func init() {

	Cmd.AddCommand(SearchCIAppTestEventsCmd)
}
