package events

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchEventsCmd = &cobra.Command{
	Use:     "search-events [payload]",
	Aliases: []string{"search"},
	Short:   "Search events",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventsListResponse
		var err error

		var body datadogV2.SearchEventsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err = api.SearchEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-events")

		cmd.Println(cmdutil.FormatJSON(res, "event"))
	},
}

func init() {
	Cmd.AddCommand(SearchEventsCmd)
}
