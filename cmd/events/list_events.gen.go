package events

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListEventsCmd = &cobra.Command{
	Use:   "list-events",
	Aliases: []string{ "list", },
	Short: "Get a list of events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err := api.ListEvents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-events: %v", err)
		}

		cmdutil.PrintJSON(res, "event")
	},
}

func init() {
	Cmd.AddCommand(ListEventsCmd)
}
