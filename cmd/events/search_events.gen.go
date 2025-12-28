package events

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchEventsCmd = &cobra.Command{
	Use:   "search-events",
	Aliases: []string{ "search", },
	Short: "Search events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err := api.SearchEvents(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchEventsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search-events: %v", err)
		}

		cmdutil.PrintJSON(res, "event")
	},
}

func init() {
	Cmd.AddCommand(SearchEventsCmd)
}
