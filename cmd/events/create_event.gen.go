package events

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateEventCmd = &cobra.Command{
	Use:   "create-event",
	Short: "Post an event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err := api.CreateEvent(client.NewContext(apiKey, appKey, site), datadogV2.EventCreateRequestPayload{})
		if err != nil {
			log.Fatalf("failed to create-event: %v", err)
		}

		cmdutil.PrintJSON(res, "events")
	},
}

func init() {
	Cmd.AddCommand(CreateEventCmd)
}
