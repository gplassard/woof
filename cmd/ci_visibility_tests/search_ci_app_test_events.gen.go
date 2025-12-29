package ci_visibility_tests

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchCIAppTestEventsCmd = &cobra.Command{
	Use:   "search-ci-app-test-events",
	
	Short: "Search tests events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.SearchCIAppTestEvents(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchCIAppTestEventsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-ci-app-test-events")

		cmd.Println(cmdutil.FormatJSON(res, "citest"))
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppTestEventsCmd)
}
