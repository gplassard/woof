package test_optimization

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchFlakyTestsCmd = &cobra.Command{
	Use: "search-flaky-tests [payload]",

	Short: "Search flaky tests",
	Long: `Search flaky tests
Documentation: https://docs.datadoghq.com/api/latest/test-optimization/#search-flaky-tests`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FlakyTestsSearchResponse
		var err error

		var body datadogV2.SearchFlakyTestsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		res, _, err = api.SearchFlakyTests(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-flaky-tests")

		cmd.Println(cmdutil.FormatJSON(res, "flaky_test"))
	},
}

func init() {
	Cmd.AddCommand(SearchFlakyTestsCmd)
}
