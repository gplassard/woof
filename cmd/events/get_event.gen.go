package events

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetEventCmd = &cobra.Command{
	Use:     "get-event [event_id]",
	Aliases: []string{"get"},
	Short:   "Get an event",
	Long: `Get an event
Documentation: https://docs.datadoghq.com/api/latest/events/#get-event`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.V2EventResponse
		var err error

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err = api.GetEvent(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-event")

		cmd.Println(cmdutil.FormatJSON(res, "events"))
	},
}

func init() {

	Cmd.AddCommand(GetEventCmd)
}
