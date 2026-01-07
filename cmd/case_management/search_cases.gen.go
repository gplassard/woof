package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchCasesCmd = &cobra.Command{
	Use: "search-cases",

	Short: "Search cases",
	Long: `Search cases
Documentation: https://docs.datadoghq.com/api/latest/case-management/#search-cases`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CasesResponse
		var err error

		optionalParams := datadogV2.NewSearchCasesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort-field") {
			val, _ := cmd.Flags().GetString("sort-field")
			optionalParams.WithSortField(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		if cmd.Flags().Changed("sort-asc") {
			val, _ := cmd.Flags().GetString("sort-asc")
			optionalParams.WithSortAsc(val)
		}

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchCases(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to search-cases")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	SearchCasesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	SearchCasesCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	SearchCasesCmd.Flags().String("sort-field", "", "Specify which field to sort")

	SearchCasesCmd.Flags().String("filter", "", "Search query")

	SearchCasesCmd.Flags().String("sort-asc", "", "Specify if order is ascending or not")

	Cmd.AddCommand(SearchCasesCmd)
}
