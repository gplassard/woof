package synthetics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchSuitesCmd = &cobra.Command{
	Use: "search-suites",

	Short: "Search test suites",
	Long: `Search test suites
Documentation: https://docs.datadoghq.com/api/latest/synthetics/#search-suites`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SyntheticsSuiteSearchResponse
		var err error

		api := datadogV2.NewSyntheticsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchSuites(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to search-suites")

		fmt.Println(cmdutil.FormatJSON(res, "suites_search"))
	},
}

func init() {

	Cmd.AddCommand(SearchSuitesCmd)
}
