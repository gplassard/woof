package events

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetEventCmd = &cobra.Command{
	Use:   "get-event [event_id]",
	Aliases: []string{ "get", },
	Short: "Get an event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err := api.GetEvent(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-event: %v", err)
		}

		cmdutil.PrintJSON(res, "events")
	},
}

func init() {
	Cmd.AddCommand(GetEventCmd)
}
