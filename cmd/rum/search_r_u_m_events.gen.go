package rum

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchRUMEventsCmd = &cobra.Command{
	Use:   "search-r-u-m-events",
	Short: "Search RUM events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.SearchRUMEvents(client.NewContext(apiKey, appKey, site), datadogV2.RUMSearchEventsRequest{})
		if err != nil {
			log.Fatalf("failed to search-r-u-m-events: %v", err)
		}

		cmdutil.PrintJSON(res, "rum")
	},
}

func init() {
	Cmd.AddCommand(SearchRUMEventsCmd)
}
