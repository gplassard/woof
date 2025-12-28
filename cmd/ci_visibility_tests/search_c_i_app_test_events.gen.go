package ci_visibility_tests

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchCIAppTestEventsCmd = &cobra.Command{
	Use:   "search_c_i_app_test_events",
	Short: "Search tests events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.SearchCIAppTestEvents(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchCIAppTestEventsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search_c_i_app_test_events: %v", err)
		}

		cmdutil.PrintJSON(res, "citest")
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppTestEventsCmd)
}
