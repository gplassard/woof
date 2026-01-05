package events

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateEventCmd = &cobra.Command{
	Use:     "create-event [payload]",
	Aliases: []string{"create"},
	Short:   "Post an event",
	Long: `Post an event
Documentation: https://docs.datadoghq.com/api/latest/events/#create-event`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventCreateResponsePayload
		var err error

		var body datadogV2.EventCreateRequestPayload
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err = api.CreateEvent(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-event")

		cmd.Println(cmdutil.FormatJSON(res, "events"))
	},
}

func init() {
	Cmd.AddCommand(CreateEventCmd)
}
