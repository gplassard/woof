package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchIncidentsCmd = &cobra.Command{
	Use:     "search-incidents [query]",
	Aliases: []string{"search"},
	Short:   "Search for incidents",
	Long: `Search for incidents
Documentation: https://docs.datadoghq.com/api/latest/incidents/#search-incidents`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.IncidentSearchResponse
		var err error

		optionalParams := datadogV2.NewSearchIncidentsOptionalParameters()

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchIncidents(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to search-incidents")

		cmd.Println(cmdutil.FormatJSON(res, "incidents_search_results"))
	},
}

func init() {

	SearchIncidentsCmd.Flags().String("include", "", "Specifies which types of related objects should be included in the response.")

	SearchIncidentsCmd.Flags().String("sort", "", "Specifies the order of returned incidents.")

	SearchIncidentsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	SearchIncidentsCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	Cmd.AddCommand(SearchIncidentsCmd)
}
