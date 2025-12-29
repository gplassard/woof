package events

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateEventCmd = &cobra.Command{
	Use:   "create-event",
	Aliases: []string{ "create", },
	Short: "Post an event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err := api.CreateEvent(client.NewContext(apiKey, appKey, site), datadogV2.EventCreateRequestPayload{})
		cmdutil.HandleError(err, "failed to create-event")

		cmd.Println(cmdutil.FormatJSON(res, "events"))
	},
}

func init() {
	Cmd.AddCommand(CreateEventCmd)
}
