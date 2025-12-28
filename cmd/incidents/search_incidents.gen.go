package incidents

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchIncidentsCmd = &cobra.Command{
	Use:   "search-incidents [query]",
	Short: "Search for incidents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		res, _, err := api.SearchIncidents(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to search-incidents: %v", err)
		}

		cmdutil.PrintJSON(res, "incidents_search_results")
	},
}

func init() {
	Cmd.AddCommand(SearchIncidentsCmd)
}
