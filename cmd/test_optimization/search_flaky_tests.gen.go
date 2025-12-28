package test_optimization

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchFlakyTestsCmd = &cobra.Command{
	Use:   "search_flaky_tests",
	Short: "Search flaky tests",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		res, _, err := api.SearchFlakyTests(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchFlakyTestsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search_flaky_tests: %v", err)
		}

		cmdutil.PrintJSON(res, "flaky_test")
	},
}

func init() {
	Cmd.AddCommand(SearchFlakyTestsCmd)
}
