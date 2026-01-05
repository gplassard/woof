package events

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListEventsCmd = &cobra.Command{
	Use:     "list-events",
	Aliases: []string{"list"},
	Short:   "Get a list of events",
	Long: `Get a list of events
Documentation: https://docs.datadoghq.com/api/latest/events/#list-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventsListResponse
		var err error

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		res, _, err = api.ListEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-events")

		cmd.Println(cmdutil.FormatJSON(res, "event"))
	},
}

func init() {
	Cmd.AddCommand(ListEventsCmd)
}
