package test_optimization

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchFlakyTestsCmd = &cobra.Command{
	Use: "search-flaky-tests",

	Short: "Search flaky tests",
	Long: `Search flaky tests
Documentation: https://docs.datadoghq.com/api/latest/test-optimization/#search-flaky-tests`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FlakyTestsSearchResponse
		var err error

		optionalParams := datadogV2.NewSearchFlakyTestsOptionalParameters()

		if cmd.Flags().Changed("payload") || cmd.Flags().Changed("payload-file") {
			err = cmdutil.UnmarshalPayload(cmd, optionalParams)
			cmdutil.HandleError(err, "failed to read payload")
		}

		api := datadogV2.NewTestOptimizationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchFlakyTests(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-flaky-tests")

		fmt.Println(cmdutil.FormatJSON(res, "flaky_test"))
	},
}

func init() {

	SearchFlakyTestsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchFlakyTestsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchFlakyTestsCmd)
}
