package test_optimization

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchFlakyTestsCmd = &cobra.Command{
	Use: "search-flaky-tests",

	Short: "Search flaky tests",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		res, _, err := api.SearchFlakyTests(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchFlakyTestsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-flaky-tests")

		cmd.Println(cmdutil.FormatJSON(res, "flaky_test"))
	},
}

func init() {
	Cmd.AddCommand(SearchFlakyTestsCmd)
}
